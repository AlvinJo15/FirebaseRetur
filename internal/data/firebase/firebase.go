package firebase

import (
	"context"
	firebaseEntity "firebase/internal/entity/firebase"
	pEntity "firebase/internal/entity/returTnIn"
	"firebase/pkg/errors"
	"firebase/pkg/firebaseclient"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type (
	// Data ...
	Data struct {
		c *firestore.Client
	}
)

// New ...
func New(fc *firebaseclient.Client) Data {
	d := Data{
		c: fc.Client,
	}
	return d
}

// InsertNewData . . .
func (d Data) InsertNewData(ctx context.Context, member firebaseEntity.Firebase) error {

	_, err := d.c.Collection("Firebase/Data/Firebase-Member").Doc(member.FirebaseID).Set(ctx, member)
	if err != nil {
		return errors.Wrap(err, "[DATA][InsertNewDataMember] Failed to save documents!")
	}

	return err
}

// TampilSemuaData ...
func (d Data) TampilSemuaData(ctx context.Context) ([]firebaseEntity.Firebase, error) {
	var (
		member  firebaseEntity.Firebase
		members []firebaseEntity.Firebase
		err     error
	)

	iter := d.c.Collection("Firebase/Data/Firebase-Member").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return members, errors.Wrap(err, "[DATA][TampilSemuaData] Failed to iterate documents!")
		}
		err = doc.DataTo(&member)
		if err != nil {
			return members, errors.Wrap(err, "[DATA][TampilSemuaData] Failed to populate struct!")
		}
		members = append(members, member)
	}
	return members, err
}

// EditData . . .
func (d Data) EditData(ctx context.Context, member firebaseEntity.Firebase) error {

	_, err := d.c.Collection("Firebase/Data/Firebase-Member").Doc(member.FirebaseID).Update(ctx, []firestore.Update{
		{Path: "FirebaseName", Value: member.FirebaseName},
		{Path: "FirebaseDOB", Value: member.FirebaseDOB},
		{Path: "FirebaseAddress", Value: member.FirebaseAddress},
		{Path: "FirebaseNoHP", Value: member.FirebaseNoHP},
		{Path: "FirebaseEmail", Value: member.FirebaseEmail},
	})
	if err != nil {
		return errors.Wrap(err, "[DATA][EditData] Failed to save documents!")
	}

	return err
}

// CariDanTampilHeaderByNoReceive ...
func (d Data) CariDanTampilHeaderByNoReceive(ctx context.Context, NoReceive string) (pEntity.HeaderRC, error) {
	var (
		header pEntity.HeaderRC
		err    error
	)
	fmt.Println("header : ", header)
	fmt.Println("no receive", NoReceive)

	doc, err := d.c.Collection("ReturTnIn/" + NoReceive + "/Header").Doc(NoReceive).Get(ctx)
	if err != nil {
		return header, errors.Wrap(err, "[DATA][CariDanTampilHeaderByNoReceive]")
	}
	err = doc.DataTo(&header)
	if err != nil {
		return header, errors.Wrap(err, "[DATA][CariDanTampilDetailByNoReceive]")
	}

	return header, err
}

// CariDanTampilDetailByNoReceive ...
func (d Data) CariDanTampilDetailByNoReceive(ctx context.Context, NoReceive string) ([]pEntity.DetailRC, error) {
	var (
		detail  pEntity.DetailRC
		details []pEntity.DetailRC
		err     error
	)

	iter := d.c.Collection("ReturTnIn/" + NoReceive + "/Detail").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return details, errors.Wrap(err, "[DATA][CariDanTampilDetailByNoReceive] Failed to iterate documents!")
		}
		err = doc.DataTo(&detail)
		if err != nil {
			return details, errors.Wrap(err, "[DATA][CariDanTampilDetailByNoReceive] Failed to populate struct!")
		}
		fmt.Println("Detail", detail)
		details = append(details, detail)
	}
	fmt.Println("detail : ", details)

	return details, err
}
