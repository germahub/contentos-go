package table

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	VoteVoterTable    uint32 = 2638131561
	VoteVoteTimeTable uint32 = 277775897
	VotePostIdTable   uint32 = 3923440502
	VoteVoterUniTable uint32 = 1965941220

	VoteVoterRow uint32 = 923013397
)

////////////// SECTION Wrap Define ///////////////
type SoVoteWrap struct {
	dba         iservices.IDatabaseRW
	mainKey     *prototype.VoterId
	watcherFlag *VoteWatcherFlag
	mKeyFlag    int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf     []byte //the buffer after the main key is encoded with prefix
	mBuf        []byte //the value after the main key is encoded
	mdFuncMap   map[string]interface{}
}

func NewSoVoteWrap(dba iservices.IDatabaseRW, key *prototype.VoterId) *SoVoteWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoVoteWrap{dba, key, nil, -1, nil, nil, nil}
	result.initWatcherFlag()
	return result
}

func (s *SoVoteWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
	if s.mKeyFlag != -1 {
		//if you have already obtained the existence status of the primary key, use it directly
		if s.mKeyFlag == 0 {
			return false
		}
		return true
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}
	if res == false {
		s.mKeyFlag = 0
	} else {
		s.mKeyFlag = 1
	}
	return res
}

func (s *SoVoteWrap) MustExist(errMsgs ...interface{}) *SoVoteWrap {
	if !s.CheckExist() {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.MustExist: %v not found", s.mainKey), errMsgs...))
	}
	return s
}

func (s *SoVoteWrap) MustNotExist(errMsgs ...interface{}) *SoVoteWrap {
	if s.CheckExist() {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.MustNotExist: %v already exists", s.mainKey), errMsgs...))
	}
	return s
}

func (s *SoVoteWrap) initWatcherFlag() {
	if s.watcherFlag == nil {
		s.watcherFlag = new(VoteWatcherFlag)
		*(s.watcherFlag) = VoteWatcherFlagOfDb(s.dba.ServiceId())
	}
}

func (s *SoVoteWrap) create(f func(tInfo *SoVote)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoVote{}
	f(val)
	if val.Voter == nil {
		val.Voter = s.mainKey
	}
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}

	buf, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	err = s.dba.Put(keyBuf, buf)
	if err != nil {
		return err
	}

	// update srt list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		return err
	}

	//update unique list
	if sucNames, err := s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delUniKeysWithNames(sucNames, val)
		s.dba.Delete(keyBuf)
		return err
	}

	s.mKeyFlag = 1

	// call watchers
	s.initWatcherFlag()
	if s.watcherFlag.AnyWatcher {
		ReportTableRecordInsert(s.dba.ServiceId(), s.dba.BranchId(), s.mainKey, val)
	}

	return nil
}

func (s *SoVoteWrap) Create(f func(tInfo *SoVote), errArgs ...interface{}) *SoVoteWrap {
	err := s.create(f)
	if err != nil {
		panic(bindErrorInfo(fmt.Errorf("SoVoteWrap.Create failed: %s", err.Error()), errArgs...))
	}
	return s
}

func (s *SoVoteWrap) getMainKeyBuf() ([]byte, error) {
	if s.mainKey == nil {
		return nil, errors.New("the main key is nil")
	}
	if s.mBuf == nil {
		var err error = nil
		s.mBuf, err = kope.Encode(s.mainKey)
		if err != nil {
			return nil, err
		}
	}
	return s.mBuf, nil
}

func (s *SoVoteWrap) modify(f func(tInfo *SoVote)) error {
	if !s.CheckExist() {
		return errors.New("the SoVote table does not exist. Please create a table first")
	}
	oriTable := s.getVote()
	if oriTable == nil {
		return errors.New("fail to get origin table SoVote")
	}

	curTable := s.getVote()
	if curTable == nil {
		return errors.New("fail to create current table SoVote")
	}
	f(curTable)

	//the main key is not support modify
	if !reflect.DeepEqual(curTable.Voter, oriTable.Voter) {
		return errors.New("primary key does not support modification")
	}

	s.initWatcherFlag()
	modifiedFields, hasWatcher, err := s.getModifiedFields(oriTable, curTable)
	if err != nil {
		return err
	}

	if modifiedFields == nil || len(modifiedFields) < 1 {
		return nil
	}

	//check whether modify sort and unique field to nil
	err = s.checkSortAndUniFieldValidity(curTable, modifiedFields)
	if err != nil {
		return err
	}

	//check unique
	err = s.handleFieldMd(FieldMdHandleTypeCheck, curTable, modifiedFields)
	if err != nil {
		return err
	}

	//delete sort and unique key
	err = s.handleFieldMd(FieldMdHandleTypeDel, oriTable, modifiedFields)
	if err != nil {
		return err
	}

	//update table
	err = s.updateVote(curTable)
	if err != nil {
		return err
	}

	//insert sort and unique key
	err = s.handleFieldMd(FieldMdHandleTypeInsert, curTable, modifiedFields)
	if err != nil {
		return err
	}

	// call watchers
	if hasWatcher {
		ReportTableRecordUpdate(s.dba.ServiceId(), s.dba.BranchId(), s.mainKey, oriTable, curTable, modifiedFields)
	}

	return nil

}

func (s *SoVoteWrap) Modify(f func(tInfo *SoVote), errArgs ...interface{}) *SoVoteWrap {
	err := s.modify(f)
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.Modify failed: %s", err.Error()), errArgs...))
	}
	return s
}

func (s *SoVoteWrap) SetPostId(p uint64, errArgs ...interface{}) *SoVoteWrap {
	err := s.modify(func(r *SoVote) {
		r.PostId = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.SetPostId( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoVoteWrap) SetUpvote(p bool, errArgs ...interface{}) *SoVoteWrap {
	err := s.modify(func(r *SoVote) {
		r.Upvote = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.SetUpvote( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoVoteWrap) SetVoteTime(p *prototype.TimePointSec, errArgs ...interface{}) *SoVoteWrap {
	err := s.modify(func(r *SoVote) {
		r.VoteTime = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.SetVoteTime( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoVoteWrap) SetWeightedVp(p string, errArgs ...interface{}) *SoVoteWrap {
	err := s.modify(func(r *SoVote) {
		r.WeightedVp = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.SetWeightedVp( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoVoteWrap) checkSortAndUniFieldValidity(curTable *SoVote, fields map[string]bool) error {
	if curTable != nil && fields != nil && len(fields) > 0 {

		if fields["VoteTime"] && curTable.VoteTime == nil {
			return errors.New("sort field VoteTime can't be modified to nil")
		}

	}
	return nil
}

//Get all the modified fields in the table
func (s *SoVoteWrap) getModifiedFields(oriTable *SoVote, curTable *SoVote) (map[string]bool, bool, error) {
	if oriTable == nil {
		return nil, false, errors.New("table info is nil, can't get modified fields")
	}
	hasWatcher := false
	fields := make(map[string]bool)

	if !reflect.DeepEqual(oriTable.PostId, curTable.PostId) {
		fields["PostId"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasPostIdWatcher
	}

	if !reflect.DeepEqual(oriTable.Upvote, curTable.Upvote) {
		fields["Upvote"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasUpvoteWatcher
	}

	if !reflect.DeepEqual(oriTable.VoteTime, curTable.VoteTime) {
		fields["VoteTime"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasVoteTimeWatcher
	}

	if !reflect.DeepEqual(oriTable.WeightedVp, curTable.WeightedVp) {
		fields["WeightedVp"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasWeightedVpWatcher
	}

	hasWatcher = hasWatcher || s.watcherFlag.WholeWatcher
	return fields, hasWatcher, nil
}

func (s *SoVoteWrap) handleFieldMd(t FieldMdHandleType, so *SoVote, fields map[string]bool) error {
	if so == nil {
		return errors.New("fail to modify empty table")
	}

	//there is no field need to modify
	if fields == nil || len(fields) < 1 {
		return nil
	}

	errStr := ""

	if fields["PostId"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldPostId(so.PostId, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "PostId")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldPostId(so.PostId, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "PostId")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldPostId(so.PostId, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "PostId")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["Upvote"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldUpvote(so.Upvote, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "Upvote")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldUpvote(so.Upvote, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "Upvote")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldUpvote(so.Upvote, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "Upvote")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["VoteTime"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldVoteTime(so.VoteTime, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "VoteTime")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldVoteTime(so.VoteTime, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "VoteTime")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldVoteTime(so.VoteTime, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "VoteTime")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["WeightedVp"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldWeightedVp(so.WeightedVp, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "WeightedVp")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldWeightedVp(so.WeightedVp, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "WeightedVp")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldWeightedVp(so.WeightedVp, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "WeightedVp")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	return nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoVoteWrap) delSortKeyVoter(sa *SoVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListVoteByVoter{}
	if sa == nil {
		val.Voter = s.GetVoter()
	} else {
		val.Voter = sa.Voter
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyVoter(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListVoteByVoter{}
	val.Voter = sa.Voter
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoVoteWrap) delSortKeyVoteTime(sa *SoVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListVoteByVoteTime{}
	if sa == nil {
		val.VoteTime = s.GetVoteTime()
		val.Voter = s.mainKey

	} else {
		val.VoteTime = sa.VoteTime
		val.Voter = sa.Voter
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyVoteTime(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListVoteByVoteTime{}
	val.Voter = sa.Voter
	val.VoteTime = sa.VoteTime
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoVoteWrap) delSortKeyPostId(sa *SoVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListVoteByPostId{}
	if sa == nil {
		val.PostId = s.GetPostId()
		val.Voter = s.mainKey

	} else {
		val.PostId = sa.PostId
		val.Voter = sa.Voter
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoVoteWrap) insertSortKeyPostId(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListVoteByPostId{}
	val.Voter = sa.Voter
	val.PostId = sa.PostId
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoVoteWrap) delAllSortKeys(br bool, val *SoVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delSortKeyVoter(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyVoteTime(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyPostId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoVoteWrap) insertAllSortKeys(val *SoVote) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoVote fail ")
	}
	if !s.insertSortKeyVoter(val) {
		return errors.New("insert sort Field Voter fail while insert table ")
	}
	if !s.insertSortKeyVoteTime(val) {
		return errors.New("insert sort Field VoteTime fail while insert table ")
	}
	if !s.insertSortKeyPostId(val) {
		return errors.New("insert sort Field PostId fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoVoteWrap) removeVote() error {
	if s.dba == nil {
		return errors.New("database is nil")
	}

	s.initWatcherFlag()

	var oldVal *SoVote
	if s.watcherFlag.AnyWatcher {
		oldVal = s.getVote()
	}

	//delete sort list key
	if res := s.delAllSortKeys(true, nil); !res {
		return errors.New("delAllSortKeys failed")
	}

	//delete unique list
	if res := s.delAllUniKeys(true, nil); !res {
		return errors.New("delAllUniKeys failed")
	}

	//delete table
	key, err := s.encodeMainKey()
	if err != nil {
		return fmt.Errorf("encodeMainKey failed: %s", err.Error())
	}
	err = s.dba.Delete(key)
	if err == nil {
		s.mKeyBuf = nil
		s.mKeyFlag = -1

		// call watchers
		if s.watcherFlag.AnyWatcher && oldVal != nil {
			ReportTableRecordDelete(s.dba.ServiceId(), s.dba.BranchId(), s.mainKey, oldVal)
		}
		return nil
	} else {
		return fmt.Errorf("database.Delete failed: %s", err.Error())
	}
}

func (s *SoVoteWrap) RemoveVote(errMsgs ...interface{}) *SoVoteWrap {
	err := s.removeVote()
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoVoteWrap.RemoveVote failed: %s", err.Error()), errMsgs...))
	}
	return s
}

////////////// SECTION Members Get/Modify ///////////////

func (s *SoVoteWrap) GetPostId() uint64 {
	res := true
	msg := &SoVote{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.PostId
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.PostId
}

func (s *SoVoteWrap) mdFieldPostId(p uint64, isCheck bool, isDel bool, isInsert bool,
	so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkPostIdIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldPostId(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldPostId(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoVoteWrap) delFieldPostId(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if !s.delSortKeyPostId(so) {
		return false
	}

	return true
}

func (s *SoVoteWrap) insertFieldPostId(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if !s.insertSortKeyPostId(so) {
		return false
	}

	return true
}

func (s *SoVoteWrap) checkPostIdIsMetMdCondition(p uint64) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) GetUpvote() bool {
	res := true
	msg := &SoVote{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Upvote
			}
		}
	}
	if !res {
		var tmpValue bool
		return tmpValue
	}
	return msg.Upvote
}

func (s *SoVoteWrap) mdFieldUpvote(p bool, isCheck bool, isDel bool, isInsert bool,
	so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkUpvoteIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldUpvote(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldUpvote(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoVoteWrap) delFieldUpvote(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) insertFieldUpvote(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) checkUpvoteIsMetMdCondition(p bool) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) GetVoteTime() *prototype.TimePointSec {
	res := true
	msg := &SoVote{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.VoteTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.VoteTime
}

func (s *SoVoteWrap) mdFieldVoteTime(p *prototype.TimePointSec, isCheck bool, isDel bool, isInsert bool,
	so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkVoteTimeIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldVoteTime(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldVoteTime(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoVoteWrap) delFieldVoteTime(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if !s.delSortKeyVoteTime(so) {
		return false
	}

	return true
}

func (s *SoVoteWrap) insertFieldVoteTime(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if !s.insertSortKeyVoteTime(so) {
		return false
	}

	return true
}

func (s *SoVoteWrap) checkVoteTimeIsMetMdCondition(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) GetVoter() *prototype.VoterId {
	res := true
	msg := &SoVote{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Voter
			}
		}
	}
	if !res {
		return nil

	}
	return msg.Voter
}

func (s *SoVoteWrap) GetWeightedVp() string {
	res := true
	msg := &SoVote{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.WeightedVp
			}
		}
	}
	if !res {
		var tmpValue string
		return tmpValue
	}
	return msg.WeightedVp
}

func (s *SoVoteWrap) mdFieldWeightedVp(p string, isCheck bool, isDel bool, isInsert bool,
	so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkWeightedVpIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldWeightedVp(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldWeightedVp(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoVoteWrap) delFieldWeightedVp(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) insertFieldWeightedVp(so *SoVote) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoVoteWrap) checkWeightedVpIsMetMdCondition(p string) bool {
	if s.dba == nil {
		return false
	}

	return true
}

////////////// SECTION List Keys ///////////////
type SVoteVoterWrap struct {
	Dba iservices.IDatabaseRW
}

func NewVoteVoterWrap(db iservices.IDatabaseRW) *SVoteVoterWrap {
	if db == nil {
		return nil
	}
	wrap := SVoteVoterWrap{Dba: db}
	return &wrap
}

func (s *SVoteVoterWrap) GetMainVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByVoter{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVoteVoterWrap) GetSubVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByVoter{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.Voter

}

func (m *SoListVoteByVoter) OpeEncode() ([]byte, error) {
	pre := VoteVoterTable
	sub := m.Voter
	if sub == nil {
		return nil, errors.New("the pro Voter is nil")
	}
	sub1 := m.Voter
	if sub1 == nil {
		return nil, errors.New("the mainkey Voter is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SVoteVoterWrap) ForEachByOrder(start *prototype.VoterId, end *prototype.VoterId, lastMainKey *prototype.VoterId,
	lastSubVal *prototype.VoterId, f func(mVal *prototype.VoterId, sVal *prototype.VoterId, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := VoteVoterTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SVoteVoteTimeWrap struct {
	Dba iservices.IDatabaseRW
}

func NewVoteVoteTimeWrap(db iservices.IDatabaseRW) *SVoteVoteTimeWrap {
	if db == nil {
		return nil
	}
	wrap := SVoteVoteTimeWrap{Dba: db}
	return &wrap
}

func (s *SVoteVoteTimeWrap) GetMainVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByVoteTime{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVoteVoteTimeWrap) GetSubVal(val []byte) *prototype.TimePointSec {
	res := &SoListVoteByVoteTime{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.VoteTime

}

func (m *SoListVoteByVoteTime) OpeEncode() ([]byte, error) {
	pre := VoteVoteTimeTable
	sub := m.VoteTime
	if sub == nil {
		return nil, errors.New("the pro VoteTime is nil")
	}
	sub1 := m.Voter
	if sub1 == nil {
		return nil, errors.New("the mainkey Voter is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SVoteVoteTimeWrap) ForEachByOrder(start *prototype.TimePointSec, end *prototype.TimePointSec, lastMainKey *prototype.VoterId,
	lastSubVal *prototype.TimePointSec, f func(mVal *prototype.VoterId, sVal *prototype.TimePointSec, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := VoteVoteTimeTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SVotePostIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewVotePostIdWrap(db iservices.IDatabaseRW) *SVotePostIdWrap {
	if db == nil {
		return nil
	}
	wrap := SVotePostIdWrap{Dba: db}
	return &wrap
}

func (s *SVotePostIdWrap) GetMainVal(val []byte) *prototype.VoterId {
	res := &SoListVoteByPostId{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Voter

}

func (s *SVotePostIdWrap) GetSubVal(val []byte) *uint64 {
	res := &SoListVoteByPostId{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return &res.PostId

}

func (m *SoListVoteByPostId) OpeEncode() ([]byte, error) {
	pre := VotePostIdTable
	sub := m.PostId

	sub1 := m.Voter
	if sub1 == nil {
		return nil, errors.New("the mainkey Voter is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SVotePostIdWrap) ForEachByOrder(start *uint64, end *uint64, lastMainKey *prototype.VoterId,
	lastSubVal *uint64, f func(mVal *prototype.VoterId, sVal *uint64, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := VotePostIdTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

/////////////// SECTION Private function ////////////////

func (s *SoVoteWrap) update(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	buf, err := proto.Marshal(sa)
	if err != nil {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	return s.dba.Put(keyBuf, buf) == nil
}

func (s *SoVoteWrap) getVote() *SoVote {
	if s.dba == nil {
		return nil
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return nil
	}
	resBuf, err := s.dba.Get(keyBuf)

	if err != nil {
		return nil
	}

	res := &SoVote{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoVoteWrap) updateVote(so *SoVote) error {
	if s.dba == nil {
		return errors.New("update fail:the db is nil")
	}

	if so == nil {
		return errors.New("update fail: the SoVote is nil")
	}

	key, err := s.encodeMainKey()
	if err != nil {
		return nil
	}

	buf, err := proto.Marshal(so)
	if err != nil {
		return err
	}

	err = s.dba.Put(key, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SoVoteWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := VoteVoterRow
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	s.mKeyBuf = kope.PackList(list)
	return s.mKeyBuf, nil
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoVoteWrap) delAllUniKeys(br bool, val *SoVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyVoter(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoVoteWrap) delUniKeysWithNames(names map[string]string, val *SoVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["Voter"]) > 0 {
		if !s.delUniKeyVoter(val) {
			res = false
		}
	}

	return res
}

func (s *SoVoteWrap) insertAllUniKeys(val *SoVote) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoVote fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyVoter(val) {
		return sucFields, errors.New("insert unique Field Voter fail while insert table ")
	}
	sucFields["Voter"] = "Voter"

	return sucFields, nil
}

func (s *SoVoteWrap) delUniKeyVoter(sa *SoVote) bool {
	if s.dba == nil {
		return false
	}
	pre := VoteVoterUniTable
	kList := []interface{}{pre}
	if sa != nil {
		if sa.Voter == nil {
			return false
		}

		sub := sa.Voter
		kList = append(kList, sub)
	} else {
		sub := s.GetVoter()
		if sub == nil {
			return true
		}

		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoVoteWrap) insertUniKeyVoter(sa *SoVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}

	pre := VoteVoterUniTable
	sub := sa.Voter
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	res, err := s.dba.Has(kBuf)
	if err == nil && res == true {
		//the unique key is already exist
		return false
	}
	val := SoUniqueVoteByVoter{}
	val.Voter = sa.Voter

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniVoteVoterWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniVoteVoterWrap(db iservices.IDatabaseRW) *UniVoteVoterWrap {
	if db == nil {
		return nil
	}
	wrap := UniVoteVoterWrap{Dba: db}
	return &wrap
}

func (s *UniVoteVoterWrap) UniQueryVoter(start *prototype.VoterId) *SoVoteWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := VoteVoterUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueVoteByVoter{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoVoteWrap(s.Dba, res.Voter)

			return wrap
		}
	}
	return nil
}

////////////// SECTION Watchers ///////////////

type VoteWatcherFlag struct {
	HasPostIdWatcher bool

	HasUpvoteWatcher bool

	HasVoteTimeWatcher bool

	HasWeightedVpWatcher bool

	WholeWatcher bool
	AnyWatcher   bool
}

var (
	VoteTable = &TableInfo{
		Name:    "Vote",
		Primary: "Voter",
		Record:  reflect.TypeOf((*SoVote)(nil)).Elem(),
	}
	VoteWatcherFlags     = make(map[uint32]VoteWatcherFlag)
	VoteWatcherFlagsLock sync.RWMutex
)

func VoteWatcherFlagOfDb(dbSvcId uint32) VoteWatcherFlag {
	VoteWatcherFlagsLock.RLock()
	defer VoteWatcherFlagsLock.RUnlock()
	return VoteWatcherFlags[dbSvcId]
}

func VoteRecordWatcherChanged(dbSvcId uint32) {
	var flag VoteWatcherFlag
	flag.WholeWatcher = HasTableRecordWatcher(dbSvcId, VoteTable.Record, "")
	flag.AnyWatcher = flag.WholeWatcher

	flag.HasPostIdWatcher = HasTableRecordWatcher(dbSvcId, VoteTable.Record, "PostId")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasPostIdWatcher

	flag.HasUpvoteWatcher = HasTableRecordWatcher(dbSvcId, VoteTable.Record, "Upvote")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasUpvoteWatcher

	flag.HasVoteTimeWatcher = HasTableRecordWatcher(dbSvcId, VoteTable.Record, "VoteTime")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasVoteTimeWatcher

	flag.HasWeightedVpWatcher = HasTableRecordWatcher(dbSvcId, VoteTable.Record, "WeightedVp")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasWeightedVpWatcher

	VoteWatcherFlagsLock.Lock()
	VoteWatcherFlags[dbSvcId] = flag
	VoteWatcherFlagsLock.Unlock()
}

////////////// SECTION Json query ///////////////

func VoteQuery(db iservices.IDatabaseRW, keyJson string) (valueJson string, err error) {
	k := new(prototype.VoterId)
	d := json.NewDecoder(bytes.NewReader([]byte(keyJson)))
	d.UseNumber()
	if err = d.Decode(k); err != nil {
		return
	}
	if v := NewSoVoteWrap(db, k).getVote(); v == nil {
		err = errors.New("not found")
	} else {
		var jbytes []byte
		if jbytes, err = json.Marshal(v); err == nil {
			valueJson = string(jbytes)
		}
	}
	return
}

func init() {
	RegisterTableWatcherChangedCallback(VoteTable.Record, VoteRecordWatcherChanged)
	RegisterTableJsonQuery("Vote", VoteQuery)
}
