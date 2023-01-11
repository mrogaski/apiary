package message_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrogaski/apiary/message"
)

func TestNew(t *testing.T) {
	t.Parallel()

	type args struct {
		r io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    *message.Message
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success",
			args:    args{r: strings.NewReader("foo")},
			want:    &message.Message{Data: []byte("foo")},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := message.New(tt.args.r)

			if tt.wantErr(t, err) {
				assert.NotEmpty(t, got.ID)
				assert.Equal(t, tt.want.Data, got.Data)
			}
		})
	}
}
