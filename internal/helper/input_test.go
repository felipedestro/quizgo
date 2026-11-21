package helper

import "testing"

func TestToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		max     int
		want    int
		wantErr bool
	}{
		{
			name:  "Número válido",
			input: "2",
			max:   4,
			want:  2,
		},
		{
			name:    "Zero inválido",
			input:   "0",
			max:     4,
			wantErr: true,
		},
		{
			name:    "Número negativo",
			input:   "-1",
			max:     4,
			wantErr: true,
		},
		{
			name:    "Número acima do limite",
			input:   "5",
			max:     4,
			wantErr: true,
		},
		{
			name:    "Texto inválido",
			input:   "abc",
			max:     4,
			wantErr: true,
		},
		{
			name:    "Entrada vazia",
			input:   "",
			max:     4,
			wantErr: true,
		},
		{
			name:  "Limite permitido",
			input: "4",
			max:   4,
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.input, tt.max)

			if (err != nil) != tt.wantErr {
				t.Fatalf(
					"ToInt(%q, %d): erro = %v, esperado erro = %v",
					tt.input, tt.max, err, tt.wantErr,
				)
			}

			if err == nil && got != tt.want {
				t.Errorf(
					"ToInt(%q, %d) = %d, esperado %d",
					tt.input, tt.max, got, tt.want,
				)
			}
		})
	}
}
