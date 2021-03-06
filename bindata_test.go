// Code generated by go-bindata. DO NOT EDIT.
// sources:
// fonts/3Dumb.ttf (142.224kB)
// fonts/ApothecaryFont.ttf (62.08kB)
// fonts/Comismsh.ttf (80.132kB)
// fonts/DENNEthree-dee.ttf (83.188kB)
// fonts/DeborahFancyDress.ttf (32.52kB)
// fonts/Flim-Flam.ttf (140.576kB)
// fonts/RitaSmith.ttf (31.24kB)
// fonts/actionj.ttf (34.944kB)
// fonts/chromohv.ttf (45.9kB)
// fonts/readme.md (162B)
// fonts/wqy-microhei.ttc (5.177MB)

package base64Captcha

import (
	"crypto/sha256"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_bindataRead(t *testing.T) {
	type args struct {
		data []byte
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bindataRead(tt.args.data, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("bindataRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bindataRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bindataFileInfo_Name(t *testing.T) {
	tests := []struct {
		name string
		fi   bindataFileInfo
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.Name(); got != tt.want {
				t.Errorf("bindataFileInfo.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bindataFileInfo_Size(t *testing.T) {
	tests := []struct {
		name string
		fi   bindataFileInfo
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.Size(); got != tt.want {
				t.Errorf("bindataFileInfo.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bindataFileInfo_Mode(t *testing.T) {
	tests := []struct {
		name string
		fi   bindataFileInfo
		want os.FileMode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.Mode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bindataFileInfo.Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bindataFileInfo_ModTime(t *testing.T) {
	tests := []struct {
		name string
		fi   bindataFileInfo
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.ModTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bindataFileInfo.ModTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bindataFileInfo_IsDir(t *testing.T) {
	tests := []struct {
		name string
		fi   bindataFileInfo
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.IsDir(); got != tt.want {
				t.Errorf("bindataFileInfo.IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bindataFileInfo_Sys(t *testing.T) {
	tests := []struct {
		name string
		fi   bindataFileInfo
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.Sys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bindataFileInfo.Sys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts3dumbTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts3dumbTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts3dumbTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts3dumbTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fonts3dumbTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fonts3dumbTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fonts3dumbTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fonts3dumbTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsApothecaryfontTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsApothecaryfontTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsApothecaryfontTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsApothecaryfontTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsApothecaryfontTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsApothecaryfontTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsApothecaryfontTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsApothecaryfontTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsComismshTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsComismshTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsComismshTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsComismshTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsComismshTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsComismshTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsComismshTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsComismshTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsDennethreeDeeTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsDennethreeDeeTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsDennethreeDeeTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsDennethreeDeeTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsDennethreeDeeTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsDennethreeDeeTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsDennethreeDeeTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsDennethreeDeeTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsDeborahfancydressTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsDeborahfancydressTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsDeborahfancydressTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsDeborahfancydressTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsDeborahfancydressTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsDeborahfancydressTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsDeborahfancydressTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsDeborahfancydressTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsFlimFlamTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsFlimFlamTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsFlimFlamTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsFlimFlamTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsFlimFlamTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsFlimFlamTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsFlimFlamTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsFlimFlamTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsRitasmithTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsRitasmithTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsRitasmithTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsRitasmithTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsRitasmithTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsRitasmithTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsRitasmithTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsRitasmithTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsActionjTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsActionjTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsActionjTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsActionjTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsActionjTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsActionjTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsActionjTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsActionjTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsChromohvTtfBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsChromohvTtfBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsChromohvTtfBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsChromohvTtfBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsChromohvTtf(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsChromohvTtf()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsChromohvTtf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsChromohvTtf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsReadmeMdBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsReadmeMdBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsReadmeMdBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsReadmeMdBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsReadmeMd(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsReadmeMd()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsReadmeMd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsReadmeMd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsWqyMicroheiTtcBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsWqyMicroheiTtcBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsWqyMicroheiTtcBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsWqyMicroheiTtcBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fontsWqyMicroheiTtc(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fontsWqyMicroheiTtc()
			if (err != nil) != tt.wantErr {
				t.Errorf("fontsWqyMicroheiTtc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fontsWqyMicroheiTtc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsset(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Asset(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Asset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Asset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssetString(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AssetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustAsset(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustAsset(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustAssetString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustAssetString(tt.args.name); got != tt.want {
				t.Errorf("MustAssetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetInfo(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssetInfo(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetDigest(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    [sha256.Size]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssetDigest(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetDigest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDigests(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string][sha256.Size]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Digests()
			if (err != nil) != tt.wantErr {
				t.Errorf("Digests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Digests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetNames(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssetNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssetDir(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestoreAsset(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RestoreAsset(tt.args.dir, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RestoreAsset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRestoreAssets(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RestoreAssets(tt.args.dir, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RestoreAssets() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_filePath(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _filePath(tt.args.dir, tt.args.name); got != tt.want {
				t.Errorf("_filePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
