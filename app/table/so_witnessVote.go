package table

import (
	"errors"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	WitnessVoteTable           = []byte("WitnessVoteTable")
	WitnessVoteVoterIdTable    = []byte("WitnessVoteVoterIdTable")
	WitnessVoteVoterIdUniTable = []byte("WitnessVoteVoterIdUniTable")
)

////////////// SECTION Wrap Define ///////////////
type SoWitnessVoteWrap struct {
	dba     iservices.IDatabaseService
	mainKey *prototype.BpVoterId
}

func NewSoWitnessVoteWrap(dba iservices.IDatabaseService, key *prototype.BpVoterId) *SoWitnessVoteWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoWitnessVoteWrap{dba, key}
	return result
}

func (s *SoWitnessVoteWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}

	return res
}

func (s *SoWitnessVoteWrap) Create(f func(tInfo *SoWitnessVote)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoWitnessVote{}
	f(val)
	if val.VoterId == nil {
		val.VoterId = s.mainKey
	}
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}
	resBuf, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	err = s.dba.Put(keyBuf, resBuf)
	if err != nil {
		return err
	}

	// update sort list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		return err
	}

	//update unique list
	if err = s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delAllUniKeys(false, val)
		s.dba.Delete(keyBuf)
		return err
	}

	return nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoWitnessVoteWrap) delSortKeyVoterId(sa *SoWitnessVote) bool {
	if s.dba == nil {
		return false
	}
	val := SoListWitnessVoteByVoterId{}
	val.VoterId = sa.VoterId
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoWitnessVoteWrap) insertSortKeyVoterId(sa *SoWitnessVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListWitnessVoteByVoterId{}
	val.VoterId = sa.VoterId
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

func (s *SoWitnessVoteWrap) delAllSortKeys(br bool, val *SoWitnessVote) bool {
	if s.dba == nil || val == nil {
		return false
	}
	res := true
	if !s.delSortKeyVoterId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoWitnessVoteWrap) insertAllSortKeys(val *SoWitnessVote) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoWitnessVote fail ")
	}
	if !s.insertSortKeyVoterId(val) {
		return errors.New("insert sort Field VoterId fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoWitnessVoteWrap) RemoveWitnessVote() bool {
	if s.dba == nil {
		return false
	}
	val := s.getWitnessVote()
	if val == nil {
		return false
	}
	//delete sort list key
	if res := s.delAllSortKeys(true, val); !res {
		return false
	}

	//delete unique list
	if res := s.delAllUniKeys(true, val); !res {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}
	return s.dba.Delete(keyBuf) == nil
}

////////////// SECTION Members Get/Modify ///////////////
func (s *SoWitnessVoteWrap) GetVoteTime() *prototype.TimePointSec {
	res := s.getWitnessVote()

	if res == nil {
		return nil

	}
	return res.VoteTime
}

func (s *SoWitnessVoteWrap) MdVoteTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	sa := s.getWitnessVote()
	if sa == nil {
		return false
	}

	sa.VoteTime = p
	if !s.update(sa) {
		return false
	}

	return true
}

func (s *SoWitnessVoteWrap) GetVoterId() *prototype.BpVoterId {
	res := s.getWitnessVote()

	if res == nil {
		return nil

	}
	return res.VoterId
}

func (s *SoWitnessVoteWrap) GetWitnessId() *prototype.BpWitnessId {
	res := s.getWitnessVote()

	if res == nil {
		return nil

	}
	return res.WitnessId
}

func (s *SoWitnessVoteWrap) MdWitnessId(p *prototype.BpWitnessId) bool {
	if s.dba == nil {
		return false
	}
	sa := s.getWitnessVote()
	if sa == nil {
		return false
	}

	sa.WitnessId = p
	if !s.update(sa) {
		return false
	}

	return true
}

////////////// SECTION List Keys ///////////////
type SWitnessVoteVoterIdWrap struct {
	Dba iservices.IDatabaseService
}

func NewWitnessVoteVoterIdWrap(db iservices.IDatabaseService) *SWitnessVoteVoterIdWrap {
	if db == nil {
		return nil
	}
	wrap := SWitnessVoteVoterIdWrap{Dba: db}
	return &wrap
}

func (s *SWitnessVoteVoterIdWrap) DelIterater(iterator iservices.IDatabaseIterator) {
	if iterator == nil || !iterator.Valid() {
		return
	}
	s.Dba.DeleteIterator(iterator)
}

func (s *SWitnessVoteVoterIdWrap) GetMainVal(iterator iservices.IDatabaseIterator) *prototype.BpVoterId {
	if iterator == nil || !iterator.Valid() {
		return nil
	}
	val, err := iterator.Value()

	if err != nil {
		return nil
	}

	res := &SoListWitnessVoteByVoterId{}
	err = proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.VoterId

}

func (s *SWitnessVoteVoterIdWrap) GetSubVal(iterator iservices.IDatabaseIterator) *prototype.BpVoterId {
	if iterator == nil || !iterator.Valid() {
		return nil
	}

	val, err := iterator.Value()

	if err != nil {
		return nil
	}
	res := &SoListWitnessVoteByVoterId{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.VoterId

}

func (m *SoListWitnessVoteByVoterId) OpeEncode() ([]byte, error) {
	pre := WitnessVoteVoterIdTable
	sub := m.VoterId
	if sub == nil {
		return nil, errors.New("the pro VoterId is nil")
	}
	sub1 := m.VoterId
	if sub1 == nil {
		return nil, errors.New("the mainkey VoterId is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query sort by order
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
func (s *SWitnessVoteVoterIdWrap) QueryListByOrder(start *prototype.BpVoterId, end *prototype.BpVoterId) iservices.IDatabaseIterator {
	if s.Dba == nil {
		return nil
	}
	pre := WitnessVoteVoterIdTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return nil
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return nil
	}
	return s.Dba.NewIterator(sBuf, eBuf)
}

/////////////// SECTION Private function ////////////////

func (s *SoWitnessVoteWrap) update(sa *SoWitnessVote) bool {
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

func (s *SoWitnessVoteWrap) getWitnessVote() *SoWitnessVote {
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

	res := &SoWitnessVote{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoWitnessVoteWrap) encodeMainKey() ([]byte, error) {
	pre := WitnessVoteTable
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	kList := []interface{}{pre, sub}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoWitnessVoteWrap) delAllUniKeys(br bool, val *SoWitnessVote) bool {
	if s.dba == nil || val == nil {
		return false
	}
	res := true
	if !s.delUniKeyVoterId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoWitnessVoteWrap) insertAllUniKeys(val *SoWitnessVote) error {
	if s.dba == nil {
		return errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert uniuqe Field fail,get the SoWitnessVote fail ")
	}
	if !s.insertUniKeyVoterId(val) {
		return errors.New("insert unique Field VoterId fail while insert table ")
	}

	return nil
}

func (s *SoWitnessVoteWrap) delUniKeyVoterId(sa *SoWitnessVote) bool {
	if s.dba == nil {
		return false
	}
	pre := WitnessVoteVoterIdUniTable
	sub := sa.VoterId
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoWitnessVoteWrap) insertUniKeyVoterId(sa *SoWitnessVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	uniWrap := UniWitnessVoteVoterIdWrap{}
	uniWrap.Dba = s.dba

	res := uniWrap.UniQueryVoterId(sa.VoterId)
	if res != nil {
		//the unique key is already exist
		return false
	}
	val := SoUniqueWitnessVoteByVoterId{}
	val.VoterId = sa.VoterId

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	pre := WitnessVoteVoterIdUniTable
	sub := sa.VoterId
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Put(kBuf, buf) == nil

}

type UniWitnessVoteVoterIdWrap struct {
	Dba iservices.IDatabaseService
}

func NewUniWitnessVoteVoterIdWrap(db iservices.IDatabaseService) *UniWitnessVoteVoterIdWrap {
	if db == nil {
		return nil
	}
	wrap := UniWitnessVoteVoterIdWrap{Dba: db}
	return &wrap
}

func (s *UniWitnessVoteVoterIdWrap) UniQueryVoterId(start *prototype.BpVoterId) *SoWitnessVoteWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := WitnessVoteVoterIdUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueWitnessVoteByVoterId{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoWitnessVoteWrap(s.Dba, res.VoterId)

			return wrap
		}
	}
	return nil
}
