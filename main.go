// SPDX-FileCopyrightText: 2025 James Pond <james@cipher.host>
//
// SPDX-License-Identifier: EUPL-1.2

package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/gousb"
	"github.com/hennedo/escpos"
)

const (
	// epsonVendorID is the USB vendor ID for Epson thermal printers.
	epsonVendorID = 0x04b8

	// epsonProductID is the USB product ID for the Epson TM-T20X-II printer.
	// epsonProductID = 0x0e27

	// pokedexBaseURL is the base URL for the Pokédex.
	pokedexBaseURL = "https://www.pokemon.com/us/pokedex/"

	// easterEggURL is a YouTube link for an easter egg.
	easterEggURL = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	// numberOfPokemons is the number of Pokémon available in the Pokédex.
	numberOfPokemons = 1025
)

var (
	epsonProductID int
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {
	var (
		task      = flag.String("task", "", "the task description")
		owner     = flag.String("owner", "", "the person responsible for the task")
		model     = flag.String("model", "TM-T20X-II", "the thermal printer model (TM-T20X-II or TM-T20III)")
		noPokedex = flag.Bool("nopokedex", false, "if set to true, no Pokedex QR code will be included")
	)

	flag.Parse()

	if *task == "" {
		fmt.Fprintf(os.Stderr, "Error: Task description is required.\n")

		os.Exit(1)
	}

	// fmt.Fprintf(os.Stdout, "Printer Model:  %s\n", *model)

	epsonProductID = 0x0e27

	switch *model {
	case "TM-T20III":
		epsonProductID = 0x0e28
		_ = epsonProductID
	case "TM-T20X-II":
		epsonProductID = 0x0e27
		_ = epsonProductID
	}

	// if *model == "TM-T20III" {
	// 	epsonProductID = 0x0e28
	// 	_ = epsonProductID
	// } else {
	// 	epsonProductID = 0x0e27
	// 	_ = epsonProductID
	// }

	// fmt.Fprintf(os.Stdout, "Printer Product ID:  %s\n", gousb.ID(epsonProductID))

	ctx := gousb.NewContext()
	defer ctx.Close()

	dev, err := ctx.OpenDeviceWithVIDPID(epsonVendorID, gousb.ID(epsonProductID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Cannot open USB device: %q\n", err)
		os.Exit(1)
	}

	if dev == nil {
		fmt.Fprintln(os.Stderr, "Error: Printer not found. Check USB connection and make sure the printer is turned on.")
		os.Exit(1)
	}

	defer dev.Close()

	// intf := ""
	// done := ""

	// if dev != nil {
	// 	intf, done, err := dev.DefaultInterface()
	// 	_, _, _ = intf, done, err
	// } else {
	// 	intf, done, err := altDev.DefaultInterface()
	// 	_, _, _ = intf, done, err
	// }
	intf, done, err := dev.DefaultInterface()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Cannot claim interface: %q\n", err)

		os.Exit(1)
	}
	defer done()

	endpoint, err := intf.OutEndpoint(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Cannot get 'out' endpoint: %q\n", err)

		os.Exit(1)
	}

	printer := escpos.New(endpoint)

	printer.Initialize()

	if *owner != "" {
		printer.Reverse(true).Bold(true).Size(1, 1).Write(strings.ToUpper(*owner))
		printer.Reverse(false).Bold(true).Size(1, 1).Write(" / " + time.Now().Format(time.DateOnly) + "\n")
	} else {
		printer.Reverse(false).Bold(true).Size(1, 1).Write(time.Now().Format(time.DateOnly) + "\n")
	}

	printer.Bold(true).Size(2, 2).Write(*task + "\n\n")

	if !*noPokedex {
		printer.QRCode(pokemonRoulette(), true, 4, escpos.QRCodeErrorCorrectionLevelH)
	}

	printer.LineFeed()

	err = printer.PrintAndCut()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Cannot print: %q\n", err)

		os.Exit(1)
	}
}

func pokemonRoulette() string {
	number := rand.IntN(numberOfPokemons)

	if number == 0 {
		return easterEggURL
	}

	return pokedexBaseURL + strconv.Itoa(number)
}
