package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddRegistration(t *testing.T) {
	type args struct {
		name string
		reg  *Registration
	}
	tests := []struct {
		name     string
		existing []args
		args     args
		wantErr  error
	}{
		{
			name:     "First registration",
			existing: []args{},
			args: args{
				name: "mongodb",
				reg:  &Registration{},
			},
			wantErr: nil,
		},
		{
			name: "Second registration",
			existing: []args{
				{
					name: "couchdb",
					reg:  &Registration{},
				},
			},
			args: args{
				name: "mongodb",
				reg:  &Registration{},
			},
			wantErr: nil,
		},
		{
			name: "Duplicate registration",
			existing: []args{
				{
					name: "mongodb",
					reg:  &Registration{},
				},
			},
			args: args{
				name: "mongodb",
				reg:  &Registration{},
			},
			wantErr: errors.New("mongodb is already registered"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry = make(map[string]*Registration) // reset registry
			for _, e := range tt.existing {
				registry[e.name] = e.reg
			}
			err := AddRegistration(tt.args.name, tt.args.reg)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
			reg, ok := registry[tt.args.name]
			assert.True(t, ok)
			assert.Equal(t, tt.args.reg, reg)
		})
	}
}

func TestGetRegistration(t *testing.T) {
	testReg := &Registration{}
	type args struct {
		name string
	}
	tests := []struct {
		name          string
		existingNames []string
		existingRegs  []*Registration
		args          args
		want          *Registration
		wantErr       error
	}{
		{
			name:          "Exist",
			existingNames: []string{"mongodb"},
			existingRegs:  []*Registration{testReg},
			args: args{
				name: "mongodb",
			},
			want:    testReg,
			wantErr: nil,
		},
		{
			name:          "Not Exist",
			existingNames: []string{},
			existingRegs:  []*Registration{},
			args: args{
				name: "mongodb",
			},
			want:    nil,
			wantErr: errors.New("cannot find mongodb repository"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry = make(map[string]*Registration) // reset registry
			for i, name := range tt.existingNames {
				registry[name] = tt.existingRegs[i]
			}
			got, err := GetRegistration(tt.args.name)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
