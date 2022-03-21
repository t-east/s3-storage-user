package mocks

import (
	entities "user/src/domains/entities"
	"user/src/usecases/port"

	"github.com/Nik-U/pbc"
)

type ContentControllerMock struct {
	CryptFactory  func() port.ContentCrypt
	OutputFactory func() port.ContentOutputPort
	InputFactory  func(
		o port.ContentOutputPort,
		cr port.ContentCrypt,
	) port.ContentInputPort
}

func NewContentCryptMock() port.ContentCrypt {
	return &ContentCryptMock{}
}

func NewContentOutputPortMock() port.ContentOutputPort {
	return &ContentOutputPortMock{}
}

type ContentRepositoryMock struct {
}

type ContentCryptMock struct {
}

type ContentOutputPortMock struct {
}

type ContentSPMock struct {
}

type ContentContractMock struct {
}

func (m *ContentCryptMock) MakeMetaData(input *entities.ContentIn) (*entities.Content, error) {
	content := &entities.Content{
		Content:     input.Content,
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: input.ContentName,
		SplitCount:  0,
	}
	return content, nil
}

func (m *ContentOutputPortMock) Render(c *entities.Content, n int) {
}

func (m *ContentOutputPortMock) RenderKey(c *entities.Key, n int) {
}

func (m *ContentOutputPortMock) RenderError(error) {
}

func CreateParamMock() (*entities.Param, *entities.Key, error) {
	params := pbc.GenerateA(uint32(160), uint32(512))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	privKey := pairing.NewZr().Rand()
	pubKey := pairing.NewG1().MulZn(g, privKey)
	p := &entities.Param{
		Pairing: params.String(),
		G:       g.Bytes(),
		U:       u.Bytes(),
	}
	k := &entities.Key{
		PubKey:  pubKey.Bytes(),
		PrivKey: privKey.Bytes(),
	}
	return p, k, nil
}

func CreateParamSample() (*entities.Param, *entities.Key, error) {
	p := &entities.Param{
		Pairing: "type a\nq 4949053928829930048246841536178180660184700616924443219116507608543628975816009400024212329964066549044583667914920506702904137063160138268320662911376143\nh 3386280112465861081465751400113171984329683467398578009076766868360154747209796407396708058660402599109872\nr 1461501637330902918203684832716283019653785059327\nexp2 160\nexp1 31\nsign1 -1\nsign0 -1",
		G:       []byte{64, 225, 26, 112, 173, 18, 15, 190, 245, 26, 185, 0, 130, 252, 71, 45, 124, 103, 158, 102, 183, 183, 132, 205, 123, 114, 114, 229, 49, 202, 10, 128, 174, 19, 110, 17, 201, 25, 174, 222, 175, 233, 42, 89, 81, 37, 107, 87, 175, 142, 194, 117, 72, 145, 107, 114, 94, 139, 36, 51, 78, 137, 101, 76, 48, 71, 111, 237, 20, 240, 0, 226, 29, 225, 52, 229, 62, 53, 64, 119, 211, 141, 20, 227, 82, 212, 126, 89, 233, 15, 48, 65, 184, 218, 140, 91, 95, 240, 55, 74, 64, 189, 208, 148, 197, 120, 222, 255, 52, 20, 177, 151, 180, 85, 86, 147, 32, 31, 170, 255, 234, 158, 121, 25, 121, 57, 90, 84},
		U:       []byte{18, 93, 40, 95, 195, 155, 153, 91, 73, 87, 39, 114, 32, 120, 27, 121, 191, 235, 149, 210, 238, 234, 254, 251, 191, 245, 91, 205, 216, 219, 50, 141, 181, 251, 243, 251, 105, 101, 53, 66, 54, 244, 50, 121, 86, 53, 53, 64, 104, 186, 56, 100, 194, 130, 1, 100, 54, 18, 28, 243, 165, 105, 193, 5, 9, 162, 239, 122, 164, 205, 235, 228, 162, 78, 103, 26, 210, 237, 29, 7, 164, 195, 24, 7, 243, 9, 147, 38, 101, 155, 56, 49, 255, 51, 179, 63, 191, 138, 5, 12, 191, 135, 140, 255, 194, 134, 177, 213, 234, 190, 83, 216, 56, 22, 218, 115, 117, 188, 174, 206, 123, 1, 94, 183, 182, 5, 159, 15},
	}
	k := &entities.Key{
		PubKey:  []byte{},
		PrivKey: []byte{},
	}
	return p, k, nil
}