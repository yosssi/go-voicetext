# go-voicetext - VoiceText Web API クライアント

[![wercker status](https://app.wercker.com/status/d481b39df7a76348daee6a5332b16475/m "wercker status")](https://app.wercker.com/project/bykey/d481b39df7a76348daee6a5332b16475)
[![Coverage Status](https://coveralls.io/repos/yosssi/go-voicetext/badge.png?branch=HEAD)](https://coveralls.io/r/yosssi/go-voicetext?branch=HEAD)
[![GoDoc](https://godoc.org/github.com/yosssi/go-voicetext?status.svg)](https://godoc.org/github.com/yosssi/go-voicetext)

## 概要

go-voicetextはGo言語向けの[VoiceText Web API](https://cloud.voicetext.jp/webapi)クライアントです。

## インストール

```sh
go get -u github.com/yosssi/go-voicetext
```

## 実装例

go-voicetextを用いた簡単な実装例を以下に記載します。

```go
package main

import (
	"fmt"
	"os"

	"github.com/yosssi/go-voicetext"
)

func main() {
	c := voicetext.NewClient(os.Getenv("VOICETEXT_API_KEY"), nil)
	result, err := c.TTS("Hello.", nil)
	if err != nil {
		panic(err)
	}

	if result.ErrMsg != nil {
		fmt.Println(result.ErrMsg)
		return
	}

	f, err := os.Create("hello.wav")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.Write(result.Sound); err != nil {
		panic(err)
	}
}
```

以下のように、`TTS`メソッドにオプションの設定を指定することができます。

```go
package main

import (
	"fmt"
	"os"

	"github.com/yosssi/go-voicetext"
)

func main() {
	c := voicetext.NewClient(os.Getenv("VOICETEXT_API_KEY"), nil)
	result, err := c.TTS("Hello.", &voicetext.TTSOptions{
		Speaker:      voicetext.SpeakerHaruka,
		Emotion:      voicetext.EmotionHappiness,
		EmotionLevel: "2",
		Pitch:        90,
		Speed:        110,
		Volume:       200,
	})
	if err != nil {
		panic(err)
	}

	if result.ErrMsg != nil {
		fmt.Println(result.ErrMsg)
		return
	}

	f, err := os.Create("hello.wav")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.Write(result.Sound); err != nil {
		panic(err)
	}
}
```

## ドキュメント

* [GoDoc](http://godoc.org/github.com/yosssi/go-voicetext)
