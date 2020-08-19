package firebase

import (
	"context"
	"firebase/internal/entity/auth"
	firebaseEntity "firebase/internal/entity/firebase"
	pEntity "firebase/internal/entity/returTnIn"
	"firebase/pkg/errors"
	"fmt"
	"strings"
	"time"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	InsertNewData(ctx context.Context, member firebaseEntity.Firebase) error
	TampilSemuaData(ctx context.Context) ([]firebaseEntity.Firebase, error)
	EditData(ctx context.Context, member firebaseEntity.Firebase) error
	CariDanTampilHeaderByNoReceive(ctx context.Context, NoReceive string) (pEntity.HeaderRC, error)
	CariDanTampilDetailByNoReceive(ctx context.Context, NoReceive string) ([]pEntity.DetailRC, error)
}

// AuthData ...
type AuthData interface {
	CheckAuth(ctx context.Context, code string) (auth.Auth, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data     Data
	authData AuthData
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data, authData AuthData) Service {
	// Assign variable dari parameter ke object
	return Service{
		data:     data,
		authData: authData,
	}
}

// Getfirebase ...
func (s Service) Getfirebase(ctx context.Context) error {

	// CHECK AUTH
	auth, err := s.authData.CheckAuth(ctx, "191")
	if err != nil {
		return errors.Wrap(err, "[SERVICE][Getfirebase]")
	}
	if auth.Error.Status == true {
		return errors.Wrap(errors.New("401 Unauthorized"), "[SERVICE][Getfirebase]")
	}
	// END CHECK AUTH

	return nil
}

// InsertNewData ...
func (s Service) InsertNewData(ctx context.Context, member firebaseEntity.Firebase) error {
	var (
		err error
	)
	err = s.data.InsertNewData(ctx, member)
	fmt.Println("masuk : ", member)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][InsertNewDataMember]")
	}

	return err
}

// TampilSemuaData ...
func (s Service) TampilSemuaData(ctx context.Context) ([]firebaseEntity.Firebase, error) {
	result, err := s.data.TampilSemuaData(ctx)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][TampilSemuaData]")
	}
	return result, err
}

// EditData ...
func (s Service) EditData(ctx context.Context, member firebaseEntity.Firebase) error {
	var err error

	err = s.data.EditData(ctx, member)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][EditData]")
	}

	return err
}

// GenerateMemberID ...
func (s Service) GenerateMemberID(ctx context.Context) (string, error) {
	var (
		urutan    int
		memberID  string
		err       error
		test      string
		tempTahun string
		tempBulan string
	)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	var now = time.Now().Format("06-01")
	var tanggal = strings.Split(now, "-")

	tempTahun = tanggal[0]
	//tempBulan = tanggal[1]
	tempBulan = "07"

	fmt.Println("Now :", now)
	fmt.Println("Tanggal :", tanggal)
	fmt.Println("tanggal1 : ", tanggal[0])
	fmt.Println("tanggal2 : ", tanggal[1])

	//urutan = 5
	//TestLooping
	for urutan = 1; urutan <= 10; {
		test = fmt.Sprintf("%04d", urutan)
		//a := strconv.Itoa(test)
		memberID = "8" + tanggal[0] + tanggal[1] + test

		//validasi reset saat ganti tahun dan bulan
		if urutan == 3 && tanggal[0] != tempTahun {
			urutan = 1
			tempTahun = tanggal[0]
		} else if urutan == 3 && tanggal[1] != tempBulan {
			urutan = 1
			tempBulan = tanggal[1]
			//fmt.Println(tempBulan)
		} else {
			urutan++
		}

		fmt.Println("memberID : ", memberID)
	}

	return memberID, err
}

// TampilAllDataByNoReceive ...
func (s Service) TampilAllDataByNoReceive(ctx context.Context, NoReceive string) (pEntity.JSONRC, error) {
	var (
		dataReceive pEntity.JSONRC
		err         error
	)
	
	//Header
	dataReceive.TranRCH, err = s.data.CariDanTampilHeaderByNoReceive(ctx, NoReceive)
	if err != nil {
		return dataReceive, errors.Wrap(err, "[SERVICE][TampilSemuaData]")
	}
	//Detail
	dataReceive.TranRCD, err = s.data.CariDanTampilDetailByNoReceive(ctx, NoReceive)
	if err != nil {
		return dataReceive, errors.Wrap(err, "[SERVICE][TampilSemuaData]")
	}

	return dataReceive, err
}
