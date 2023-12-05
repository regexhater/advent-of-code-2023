package main

import (
	"testing"
)

func Test_isAdjacentToSymbol(t *testing.T) {
	type args struct {
		schematics []string
		x          int
		startY     int
		endY       int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "adjacent left",
			args: args{
				schematics: []string{"#412.."},
				x:          0,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent left top diagonal",
			args: args{
				schematics: []string{"#....", ".321.."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent top ",
			args: args{
				schematics: []string{".#...", ".321.."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent right top diagonal",
			args: args{
				schematics: []string{"..#..", ".321.."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent right",
			args: args{
				schematics: []string{".412#."},
				x:          0,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent right bottom diagonal",
			args: args{
				schematics: []string{".....", ".321..", "..#.."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent bottom",
			args: args{
				schematics: []string{".....", ".321..", ".#..."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "adjacent left bottom diagonal",
			args: args{
				schematics: []string{".....", ".321..", "#...."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: true,
		},
		{
			name: "no adjacent dots",
			args: args{
				schematics: []string{".....", ".321..", "....."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: false,
		},
		{
			name: "no adjacent other digits",
			args: args{
				schematics: []string{"......", ".321..", "12...."},
				x:          1,
				startY:     1,
				endY:       3,
			},
			want: false,
		},
		{
			name: "no adjacent left edge",
			args: args{
				schematics: []string{"......", "321...", "......"},
				x:          1,
				startY:     0,
				endY:       2,
			},
			want: false,
		},
		{
			name: "no adjacent top edge",
			args: args{
				schematics: []string{".321..", "......", "......"},
				x:          0,
				startY:     1,
				endY:       3,
			},
			want: false,
		},
		{
			name: "no adjacent right edge",
			args: args{
				schematics: []string{"......", "...321", "......"},
				x:          1,
				startY:     3,
				endY:       5,
			},
			want: false,
		},
		{
			name: "no adjacent bottom edge",
			args: args{
				schematics: []string{"......", "......", ".321.."},
				x:          2,
				startY:     1,
				endY:       3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isAdjacentToSymbol(
					tt.args.schematics,
					tt.args.x,
					tt.args.startY,
					tt.args.endY,
				); got != tt.want {
					t.Errorf("isAdjacentToSymbol() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
