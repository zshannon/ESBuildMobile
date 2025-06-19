package esbuildmobile

import (
	"log"
	"testing"
)

func TestTransform(t *testing.T) {
	jsx := `
	// import * as React from 'react'
	// import * as ReactDOM from 'react-dom'
        (value => {
            const host = createElement()
            render(<text>{value}</text>, host)
            return host.element
		}
        )
	`

	code, err := TransformJSX(jsx, &TransformOptions{JSXFactory: "banana", JSXFragment: "pizza"})

	if err != nil {
		t.Fatal("could not format jsx: ", err)
	}

	log.Println("transformed string: \n", code)
}
