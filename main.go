package main

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

// Função para realizar a cifra XOR (serve para criptografar e descriptografar)
func xorCipher(input, key string) string {
	output := make([]rune, len(input))
	for i, char := range input {
		output[i] = char ^ rune(key[i%len(key)])
	}
	return string(output)
}

// Componente da interface gráfica
type EncryptComponent struct {
	vecty.Core

	inputText       string // Texto a ser criptografado
	keyText         string // Chave para criptografia/descriptografia
	encryptedText   string // Texto criptografado gerado
	decryptedText   string // Resultado da descriptografia
	toDecryptedText string // Resultado da descriptografia
	toKeyText       string // Chave para criptografia/descriptografia
}

// Função que será chamada quando o botão de criptografia for clicado
func (c *EncryptComponent) onEncrypt(event *vecty.Event) {

	// Concatenar o sal com a chave
	combinedKey := c.keyText
	// Criptografar o texto
	c.encryptedText = xorCipher(c.inputText, combinedKey)

	// Salvar o sal e atualizar a interface gráfica
	vecty.Rerender(c)
}

// Função que será chamada quando o botão de descriptografia for clicado (usando o texto gerado no mesmo ciclo)
func (c *EncryptComponent) onDecrypt(event *vecty.Event) {
	// Combinar o sal armazenado com a chave para descriptografar o texto
	combinedKey := c.toKeyText
	c.decryptedText = xorCipher(c.toDecryptedText, combinedKey)

	// Atualizar a interface gráfica com o texto descriptografado
	vecty.Rerender(c)
}

// Função para renderizar a interface gráfica
func (c *EncryptComponent) Render() vecty.ComponentOrHTML {

	return elem.Body(
		// Entrada para o texto a ser criptografado

		elem.Div(
			elem.Input(
				vecty.Markup(
					vecty.Property("placeholder", "Digite o texto para criptografar"),
					event.Input(func(e *vecty.Event) {
						c.inputText = e.Target.Get("value").String()
					}),
				),
			),
		),
		// Entrada para a chave de criptografia
		elem.Div(
			elem.Input(
				vecty.Markup(
					vecty.Property("placeholder", "Digite a chave"),
					event.Input(func(e *vecty.Event) {
						c.keyText = e.Target.Get("value").String()
					}),
				),
			),
		),
		// Botão para criptografar o texto
		elem.Button(
			vecty.Text("Criptografar"),
			vecty.Markup(
				event.Click(c.onEncrypt),
			),
		),
		// Exibição do texto criptografado gerado
		elem.Div(
			vecty.Text(fmt.Sprintf("Texto criptografado: %s", c.encryptedText)),
		),

		elem.Div(
			elem.Input(
				vecty.Markup(
					vecty.Property("placeholder", "Digite o texto para Descriptografar"),
					event.Input(func(e *vecty.Event) {
						c.toDecryptedText = e.Target.Get("value").String()
					}),
				),
			),
		),
		// Entrada para a chave de criptografia
		elem.Div(
			elem.Input(
				vecty.Markup(
					vecty.Property("placeholder", "Digite a chave"),
					event.Input(func(e *vecty.Event) {
						c.toKeyText = e.Target.Get("value").String()
					}),
				),
			),
		),
		// Botão para descriptografar o texto gerado
		elem.Button(
			vecty.Text("Descriptografar o texto gerado"),
			vecty.Markup(
				event.Click(c.onDecrypt),
			),
		),
		// Exibição do resultado da descriptografia
		elem.Div(
			vecty.Text(fmt.Sprintf("Texto descriptografado: %s", c.decryptedText)),
		),
	)
}

func main() {
	// Configurar o componente na interface
	fmt.Println("oiS")
	vecty.SetTitle("Sistema de Criptografia e Descriptografia com Vecty")
	vecty.RenderBody(&EncryptComponent{})
}
