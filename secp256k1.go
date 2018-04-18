package secp256k1

// #include "c-secp256k1/include/secp256k1.h"
// #cgo LDFLAGS: ${SRCDIR}/c-secp256k1/.libs/libsecp256k1.a -lgmp
import "C"
import "unsafe"

const hsize int = 32

func Start() {
	C.secp256k1_start(C.SECP256K1_START_VERIFY | C.SECP256K1_START_SIGN)
}

func Seckey_verify(seckey [32]byte) bool {
	success := C.secp256k1_ec_seckey_verify(sliceToUchar(seckey[:]))
	return checkExitCode(success)
}

func Pubkey_verify(pubkey []byte) bool {
	success := C.secp256k1_ec_pubkey_verify(sliceToUchar(pubkey), C.int(len(pubkey)))
	return checkExitCode(success)
}

func Pubkey_create(seckey [32]byte, compressed bool) ([]byte, bool) {
	comp := C.int(0)
	bufsize := 65
	if compressed {
		comp = 1
		bufsize = 33
	}
	pubkey := make([]C.uchar, bufsize)
	pubkeylen := C.int(0)
	success := C.secp256k1_ec_pubkey_create(&pubkey[0],
		&pubkeylen,
		sliceToUchar(seckey[:]),
		comp)
	return C.GoBytes(unsafe.Pointer(&pubkey[0]), pubkeylen), checkExitCode(success)
}

func Sign(msgHash [hsize]byte, seckey [32]byte, nonce *[32]byte) ([]byte, bool) {
	var sig [72]C.uchar
	siglen := C.int(len(sig))
	exitCode := C.secp256k1_ecdsa_sign(sliceToUchar(msgHash[:]),
		&sig[0],
		&siglen,
		sliceToUchar(seckey[:]),
		nil,
		unsafe.Pointer(nonce))
	return C.GoBytes(unsafe.Pointer(&sig[0]), siglen), checkExitCode(exitCode)
}

func Stop() {
	C.secp256k1_stop()
}

func Verify(msgHash [hsize]byte, sig []byte, pubkey []byte) bool {
	exitCode := C.secp256k1_ecdsa_verify(sliceToUchar(msgHash[:]),
		sliceToUchar(sig),
		C.int(len(sig)),
		sliceToUchar(pubkey),
		C.int(len(pubkey)))
	return checkExitCode(exitCode)
}

func checkExitCode(success C.int) bool {
	return success == 1
}

func sliceToUchar(slice []byte) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&slice[0]))
}
