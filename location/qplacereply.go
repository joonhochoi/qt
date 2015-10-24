package location

//#include "qplacereply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlaceReply struct {
	core.QObject
}

type QPlaceReplyITF interface {
	core.QObjectITF
	QPlaceReplyPTR() *QPlaceReply
}

func PointerFromQPlaceReply(ptr QPlaceReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReplyPTR().Pointer()
	}
	return nil
}

func QPlaceReplyFromPointer(ptr unsafe.Pointer) *QPlaceReply {
	var n = new(QPlaceReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlaceReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlaceReply) QPlaceReplyPTR() *QPlaceReply {
	return ptr
}

//QPlaceReply::Error
type QPlaceReply__Error int

var (
	QPlaceReply__NoError                   = QPlaceReply__Error(0)
	QPlaceReply__PlaceDoesNotExistError    = QPlaceReply__Error(1)
	QPlaceReply__CategoryDoesNotExistError = QPlaceReply__Error(2)
	QPlaceReply__CommunicationError        = QPlaceReply__Error(3)
	QPlaceReply__ParseError                = QPlaceReply__Error(4)
	QPlaceReply__PermissionsError          = QPlaceReply__Error(5)
	QPlaceReply__UnsupportedError          = QPlaceReply__Error(6)
	QPlaceReply__BadArgumentError          = QPlaceReply__Error(7)
	QPlaceReply__CancelError               = QPlaceReply__Error(8)
	QPlaceReply__UnknownError              = QPlaceReply__Error(9)
)

//QPlaceReply::Type
type QPlaceReply__Type int

var (
	QPlaceReply__Reply                 = QPlaceReply__Type(0)
	QPlaceReply__DetailsReply          = QPlaceReply__Type(1)
	QPlaceReply__SearchReply           = QPlaceReply__Type(2)
	QPlaceReply__SearchSuggestionReply = QPlaceReply__Type(3)
	QPlaceReply__ContentReply          = QPlaceReply__Type(4)
	QPlaceReply__IdReply               = QPlaceReply__Type(5)
	QPlaceReply__MatchReply            = QPlaceReply__Type(6)
)