package main

import (
	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	// TODO: answer here
	te.UndoStack.Push(ch)
}

func (te *TextEditor) Read() []rune {
	// TODO: answer here
	if te.UndoStack.IsEmpty() {
		return nil
	}
	return te.UndoStack.Data
}

func (te *TextEditor) Undo() {
	// TODO: answer here
	recent, _ := te.UndoStack.Peek()
	if len(te.UndoStack.Data) == 0 {
		return
	}
	te.RedoStack.Push(recent)
	te.UndoStack.Pop()
}

func (te *TextEditor) Redo() {
	// TODO: answer here
	recent, _ := te.RedoStack.Peek()
	if len(te.RedoStack.Data) == 0 {
		return
	}
	te.UndoStack.Push(recent)
	te.RedoStack.Pop()
}
