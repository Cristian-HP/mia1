package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//esto para los struct y las clases dado que la mierda no sirve para algomas
type token struct {
	lexema string
	tipo   int64
	codigo int64
}

type atributo struct {
	name  string
	valor string
}
type comando struct {
	name    string
	codigo  int64
	lisAtri []atributo
}

type mBR struct {
	Mbrtamano        int64
	Mbrfechacreacion [19]byte
	Mbrdisksignature int64
	Particiones      [4]pARTICION
}
type pARTICION struct {
	Partstatus byte
	Parttype   byte
	Partfit    byte
	Partstart  int64
	Partsize   int64
	Partname   [16]byte
	Formateada int64
}

type eBR struct {
	Partstatus byte
	Partfit    byte
	Partstart  int64
	Partsize   int64
	Partnext   int64
	Partname   [16]byte
}

type nodom struct {
	Path   string
	Name   string
	Letra  byte
	Numero int64
}

//este apartado es para las variable globales para teneer por como siem
var listoken []token
var listcomand []comando
var tipo int64
var listamontada = make(map[string]nodom)

// fin apartado de variables de globales

func main() {
	//leercomando("fdisk -sizE->10 -UniT->M -path->\"/home/archivos/fase 2/D1.dsk\" -type->P -fit->FF -name->\"PRI1\" fdisk -path->\"/home/archivos/fase 2/D1.dsk\" -sizE->10000 -fit->BF -name->\"PRI2\" \n FdisK -path->\"/home/archivos/fase 2/D1.dsk\" -type->E -name->\"EXT\" -sizE->51200 \n fdisk -type->L -sizE->5120 -Unit->K -path->\"/home/archivos/fase 2/D1.dsk\" -name->\"LOG1\"")
	fmt.Println(":::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: ")
	fmt.Println(":::::::::::::::::::::::::::::::: SISTEMA DE ARCHIVOS 2.0 ::::::::::::::::::::::::::::::: ")
	fmt.Println("::::::::::::::::::::::::::::: Cristian            201603198 :::::::::::::::::::::::::::: ")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("Por favor Ingrese el comando:")
	for {
		print("$$  ")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		eleccion := strings.TrimRight(entrada, "\r\n")
		if eleccion == "0" {
			break
		} else {
			//leer el comando
			listoken = nil
			listcomand = nil
			leercomando(eleccion)
			eleccion = ""
		}
	}
}

//esto es por la mierda del lenguaje propuesto iguales los mierdas
func crearDirectorioF(path string) {
	pa := path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(pa, 0777)
		if err != nil {
			fmt.Println("fallo la mierda")
		}
	}
}

func mKdisk(size int64, unit byte, path string, name string) {

	crearDirectorioF(path)
	aux1 := path + name
	file, err := os.OpenFile(aux1, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil { //validar que no sea nulo.
		log.Fatal(err)
	}
	if err == nil {
		if unit == 'M' {
			size = size * 1024 * 1024
		} else if unit == 'K' {
			size = size * 1024
		}
		/*var i int64
		d2 := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i = 0; i < (size / 1024); i++ {
			//fmt.Fprintf(file, "%c", buffer1)
			file.Write(d2)
		}*/

		var otro int8 = 0

		ss := &otro
		//fmt.Println(unsafe.Sizeof(otro))
		var binario bytes.Buffer
		binary.Write(&binario, binary.BigEndian, ss)
		escribirBytes(file, binario.Bytes())
		//Nos posicionamos en el byte 1023 (primera posicion es 0)
		file.Seek(size, 0) // segundo parametro: 0, 1, 2.     0 -> Inicio, 1-> desde donde esta el puntero, 2 -> Del fin para atras

		//Escribimos un 0 al final del archivo.
		var binario22 bytes.Buffer
		binary.Write(&binario22, binary.BigEndian, ss)
		escribirBytes(file, binario22.Bytes())

		// Crear el MBR
		file.Seek(0, 0)
		mbr := mBR{Mbrtamano: size}
		//memset(&mbr, 0, sizeof(MBR))
		mbr.Mbrtamano = size
		mbr.Mbrdisksignature = rand.Int63n(100)
		current := time.Now()
		var fecha string = current.Format("2006-01-02 15:04:05")
		copy(mbr.Mbrfechacreacion[:], fecha)
		//mbr.fit_disk = fitt
		for p := 0; p < 4; p++ {
			mbr.Particiones[p].Partstatus = 'N'
			mbr.Particiones[p].Partsize = 0
			mbr.Particiones[p].Parttype = 'N'
			mbr.Particiones[p].Partfit = 'N'
			mbr.Particiones[p].Partstart = -1
			mbr.Particiones[p].Formateada = -1
			copy(mbr.Particiones[p].Partname[:], "")
		}
		//fmt.Println(mbr)
		//fwrite(&mbr, sizeof(MBR), 1, file)
		s := &mbr
		var binario2 bytes.Buffer
		binary.Write(&binario2, binary.BigEndian, s)
		escribirBytes(file, binario2.Bytes())
		fmt.Println("Disco creado satisfactoriamente ")

	} else {
		fmt.Println("Error al intentar crear el disco indicado")
	}
}

func analisis(entrada string) {
	var estado int32
	var auxlexema string
	for k := 0; k < len(entrada); {
		switch estado {
		case 0:
			{
				if entrada[k] == '"' {
					estado = 7
					k++

				} else if entrada[k] == '/' {
					auxlexema += string(entrada[k])
					estado = 8
					k++
				} else if entrada[k] == '#' {
					estado = 30
					k++
				} else if entrada[k] == '*' {
					estado = 0
					k++
					k++
				} else if unicode.IsDigit(rune(entrada[k])) {
					if len(entrada) > k+1 {
						estado = 3
						auxlexema += string(entrada[k])
						k++
					} else if len(entrada) <= k+1 {
						listoken = append(listoken, token{lexema: string(entrada[k]), tipo: 3, codigo: 10})
						k++
						auxlexema = ""
						estado = 0
					}
				} else if unicode.IsLetter(rune(entrada[k])) {
					if len(entrada) > k+1 {
						estado = 1
						auxlexema += string(entrada[k])
						k++
					} else if len(entrada) <= k+1 {
						listoken = append(listoken, token{lexema: string(entrada[k]), tipo: 3, codigo: 24})
						k++
						auxlexema = ""
						estado = 0
					}

				} else if entrada[k] == '-' {
					estado = 4
					k++
				} else {
					estado = 0
					k++
				}
				break
			}
		case 1:
			{
				if unicode.IsLetter(rune(entrada[k])) {
					estado = 1
					auxlexema += string(entrada[k])
					k++
				} else if unicode.IsDigit(rune(entrada[k])) {
					estado = 2
					auxlexema += string(rune(entrada[k]))
					k++
				} else if entrada[k] == '_' {
					estado = 2
					auxlexema += string(rune(entrada[k]))
					k++
				} else if entrada[k] == '.' {
					estado = 11
					auxlexema += string(rune(entrada[k]))
					k++
				} else {
					tipotoken(auxlexema)
					listoken = append(listoken, token{lexema: auxlexema, tipo: tipo, codigo: 7})
					auxlexema = ""
					estado = 0
				}
				break
			}
		case 2:
			{
				if unicode.IsLetter(rune(entrada[k])) {
					estado = 2
					auxlexema += string(entrada[k])
					k++
				} else if unicode.IsDigit(rune(entrada[k])) {
					estado = 2
					auxlexema += string(rune(entrada[k]))
					k++
				} else if entrada[k] == '_' {
					estado = 2
					auxlexema += string(rune(entrada[k]))
					k++
				} else {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 7})
					auxlexema = ""
					estado = 0
				}
				break
			}
		case 3:
			{
				if unicode.IsDigit(rune(entrada[k])) {
					estado = 3
					auxlexema += string(entrada[k])
					k++
				} else {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 8})
					auxlexema = ""
					estado = 0
				}
				break
			}
		case 4:
			{
				if unicode.IsLetter(rune(entrada[k])) {
					estado = 5
					auxlexema += string(entrada[k])
					k++
				} else if entrada[k] == '>' {
					estado = 0
					auxlexema = ""
					k++
				} else {
					estado = 0
					auxlexema = ""
				}
				break
			}
		case 5:
			{
				if unicode.IsLetter(rune(entrada[k])) {
					estado = 5
					auxlexema += string(entrada[k])
					k++
				} else if unicode.IsDigit(rune(entrada[k])) {
					estado = 6
					auxlexema += string(entrada[k])
					k++
				} else {
					estado = 0
					listoken = append(listoken, token{lexema: auxlexema, tipo: 2, codigo: 8})
					auxlexema = ""
				}
				break
			}
		case 6:
			{
				if unicode.IsDigit(rune(entrada[k])) {
					estado = 6
					auxlexema += string(entrada[k])
					k++
				} else {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 2, codigo: 8})
					estado = 0
					auxlexema = ""
				}
				break
			}
		case 7:
			{
				if entrada[k] == '"' {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 8})
					estado = 0
					auxlexema = ""
					k++
				} else {
					auxlexema += string(entrada[k])
					k++
				}
				break
			}
		case 8:
			{
				if unicode.IsLetter(rune(entrada[k])) {
					estado = 9
					auxlexema += string(entrada[k])
					k++
				} else {
					estado = 0
					auxlexema = ""
				}
				break
			}
		case 9:
			{
				if unicode.IsLetter(rune(entrada[k])) || unicode.IsDigit(rune(entrada[k])) {
					auxlexema += string(entrada[k])
					estado = 9
					k++
				} else if entrada[k] == '/' {
					auxlexema += string(entrada[k])
					estado = 10
					k++
				} else if entrada[k] == '.' {
					estado = 12
					auxlexema += string(entrada[k])
					k++
				} else {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 8})
					estado = 0
					auxlexema = ""
				}
				break
			}
		case 10:
			{
				if unicode.IsLetter(rune(entrada[k])) || unicode.IsDigit(rune(entrada[k])) {
					auxlexema += string(entrada[k])
					estado = 10
					k++
				} else if entrada[k] == '/' {
					auxlexema += string(entrada[k])
					estado = 10
					k++
				} else if entrada[k] == '.' {
					estado = 12
					auxlexema += string(entrada[k])
					k++
				} else {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 7})
					estado = 0
					auxlexema = ""
				}
				break
			}
		case 11:
			{
				if entrada[k] == 'd' && entrada[k+1] == 's' && entrada[k+2] == 'k' {
					auxlexema += "dsk"
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 5})
					estado = 0
					auxlexema = ""
					k += 3
				} else {
					fmt.Println("NO se acepta extensiones diferente de dsk")
					k++
					estado = 0
				}
				break
			}
		case 12:
			{
				if unicode.IsLetter(rune(entrada[k])) || unicode.IsDigit(rune(entrada[k])) {
					auxlexema += string(entrada[k])
					k++
					estado = 12
				} else {
					listoken = append(listoken, token{lexema: auxlexema, tipo: 3, codigo: 5})
					estado = 0
					auxlexema = ""
				}
			}

		case 30:
			{
				if entrada[k] == '\n' {
					estado = 0
					k++
				} else {
					estado = 30
					k++
				}
			}
		}
	}
}

func tipotoken(palabra string) {
	if strings.EqualFold(palabra, "exec") {
		tipo = 1
	} else if strings.EqualFold(palabra, "pause") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mkdisk") {
		tipo = 1
	} else if strings.EqualFold(palabra, "rmdisk") {
		tipo = 1
	} else if strings.EqualFold(palabra, "fdisk") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mount") {
		tipo = 1
	} else if strings.EqualFold(palabra, "unmount") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mkfs") {
		tipo = 1
	} else if strings.EqualFold(palabra, "login") {
		tipo = 1
	} else if strings.EqualFold(palabra, "logout") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mkgrp") {
		tipo = 1
	} else if strings.EqualFold(palabra, "rmgrp") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mkusr") {
		tipo = 1
	} else if strings.EqualFold(palabra, "rmusr") {
		tipo = 1
	} else if strings.EqualFold(palabra, "chmod") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mkfile") {
		tipo = 1
	} else if strings.EqualFold(palabra, "cat") {
		tipo = 1
	} else if strings.EqualFold(palabra, "rm") {
		tipo = 1
	} else if strings.EqualFold(palabra, "edit") {
		tipo = 1
	} else if strings.EqualFold(palabra, "ren") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mkdir") {
		tipo = 1
	} else if strings.EqualFold(palabra, "cp") {
		tipo = 1
	} else if strings.EqualFold(palabra, "mv") {
		tipo = 1
	} else if strings.EqualFold(palabra, "find") {
		tipo = 1
	} else if strings.EqualFold(palabra, "chown") {
		tipo = 1
	} else if strings.EqualFold(palabra, "chgrp") {
		tipo = 1
	} else if strings.EqualFold(palabra, "loss") {
		tipo = 1
	} else if strings.EqualFold(palabra, "Recovery") {
		tipo = 1
	} else if strings.EqualFold(palabra, "rep") {
		tipo = 1
	} else {
		tipo = 3 //esto para que sea los atri
	}
}

func leercomando(entrada string) {
	//obtengo la lista de valores
	analisis(entrada) //debe existir alguna entrada
	armarcomando()
	ejecutarcomando()
}

func armarcomando() {
	var lisatri []atributo
	var namecomando string
	var nameatri string
	var numcomand int = 1

	for i, s := range listoken {
		if s.tipo == 1 {
			if i > numcomand {
				listcomand = append(listcomand, comando{name: namecomando, codigo: 1, lisAtri: lisatri})
				numcomand = i
				lisatri = nil
			}
			namecomando = s.lexema
		} else if s.tipo == 2 {
			nameatri = s.lexema
		} else if s.tipo == 3 {
			lisatri = append(lisatri, atributo{name: nameatri, valor: s.lexema})
		}
	}
	listcomand = append(listcomand, comando{name: namecomando, codigo: 1, lisAtri: lisatri})
}

func ejecutarcomando() {
	//fmt.Println(len(listcomand))
	for _, com := range listcomand {
		if strings.EqualFold(com.name, "exec") {
			ejecutarEXEC(com.lisAtri)
		} else if strings.EqualFold(com.name, "mkdisk") {
			ejecutarMKDISk(com.lisAtri)
		} else if strings.EqualFold(com.name, "rmdisk") {
			ejecutarRMDISK(com.lisAtri)
		} else if strings.EqualFold(com.name, "fdisk") {
			//fmt.Println("estre al fdisk")
			ejecutarFDISK(com.lisAtri)
		} else if strings.EqualFold(com.name, "mount") {
			//fmt.Println("esto es el de ejecutat")
			//fmt.Println(com.name)
			ejecutarMOUNT(com.lisAtri)
		} else if strings.EqualFold(com.name, "UNMOUNT") {
			ejecutarUNMOUNT(com.lisAtri)
		} else if strings.EqualFold(com.name, "rep") {
			ejecutarREP(com.lisAtri)
		}
	}
}

func ejecutarEXEC(lisatributo []atributo) {
	if len(lisatributo) == 1 {
		if strings.EqualFold(lisatributo[0].name, "path") {
			file, err := ioutil.ReadFile(lisatributo[0].valor) // just pass the file name
			if err != nil {
				fmt.Println("Error en archivo indicado no existe favor verificar")
				fmt.Print(err)
			} else {
				entrada := string(file)
				/*analisis(entrada)
				fmt.Println(len(listoken))
				armarcomando()
				fmt.Println(listcomand)
				fmt.Println(len(listcomand))*/
				listcomand = nil
				listoken = nil
				lisatributo = nil
				leercomando(entrada)
			}

		} else {
			fmt.Println("el comando exec no admite el parametro -> " + lisatributo[0].name)
		}

	} else {
		fmt.Println("Comando exec solo puede aceptar un parametro")
	}
}

func ejecutarMKDISk(lisatributo []atributo) {
	var size int64 = 0
	var path string = " "
	var name string = " "
	var unit byte = 'M'
	var obli int32 = 0
	var error bool = false

	for _, param := range lisatributo {
		if strings.EqualFold(param.name, "size") {
			size1, _ := strconv.ParseInt(param.valor, 10, 64)
			size = size1
			obli++
			if size <= 0 {
				fmt.Println("ERROR: El parametro size debe ser mayor a Cer0")
				error = true
				break
			}
		} else if strings.EqualFold(param.name, "path") {
			path = param.valor
			obli++
		} else if strings.EqualFold(param.name, "name") {
			name = param.valor
			obli++
		} else if strings.EqualFold(param.name, "unit") {
			unit = param.valor[0]
			if unit == 'K' || unit == 'k' {
				unit = 'K'
			} else if unit == 'M' || unit == 'm' {
				unit = 'M'
			} else {
				fmt.Print("ERROR: el paramtro unit no acepta el valor de ")
				fmt.Println(unit)
				error = true
				break
			}
		} else {
			fmt.Println("ERROR: Parametro  NO permitido para el comando MKDISK -> " + param.name)
			error = true
			break
		}

	}

	if !error {
		if obli == 3 {
			//llamamos al mkdisk
			mKdisk(size, unit, path, name)
		}
	}

}

func ejecutarRMDISK(lisatributo []atributo) {
	var path string = ""
	if lisatributo != nil {
		if strings.EqualFold(lisatributo[0].name, "path") {
			path = lisatributo[0].valor
			rmdisk(path)
		} else {
			fmt.Println("ERROR: el comando rmdisk solo acepta el parameto path")
		}
	}
}

func ejecutarFDISK(lisatributo []atributo) {
	var size1 int64 = 0
	var add1 int64 = 0
	var unit1 byte = 'N'
	var fit1 byte = 'N'
	var error bool = false
	var path string = ""
	var name string = ""
	var type1 byte = 'P'
	var tdelete string = ""
	var bandelete bool = false
	var bandadd bool = false
	var bandsize bool = false
	var bandpath bool = false
	var bandname bool = false
	var bandtype bool = false
	for _, param := range lisatributo {
		if strings.EqualFold(param.name, "size") {
			sizeP, _ := strconv.ParseInt(param.valor, 10, 64)
			size1 = sizeP
			bandsize = true
			if size1 <= 0 {
				fmt.Println("ERROR: El parametro size debe ser mayor a cer0")
				error = true
				break
			}
		} else if strings.EqualFold(param.name, "unit") {
			if param.valor[0] == 'B' || param.valor[0] == 'b' {
				unit1 = 'B'
			} else if param.valor[0] == 'K' || param.valor[0] == 'k' {
				unit1 = 'K'
			} else if param.valor[0] == 'M' || param.valor[0] == 'm' {
				unit1 = 'M'
			} else {
				fmt.Println("ERROR: El parametro unit no admite este valor -> " + param.valor)
				error = true
				break
			}
		} else if strings.EqualFold(param.name, "path") {
			path = param.valor
			bandpath = true
		} else if strings.EqualFold(param.name, "type") {
			bandtype = true
			if param.valor[0] == 'P' || param.valor[0] == 'p' {
				type1 = 'P'
			} else if param.valor[0] == 'E' || param.valor[0] == 'e' {
				type1 = 'E'
			} else if param.valor[0] == 'L' || param.valor[0] == 'l' {
				type1 = 'L'
			} else {
				fmt.Println("ERROR: El parametro type no admite este valor -> " + param.valor)
				error = true
				break
			}
		} else if strings.EqualFold(param.name, "fit") {
			if strings.EqualFold(param.valor, "BF") {
				fit1 = 'B'
			} else if strings.EqualFold(param.valor, "FF") {
				fit1 = 'F'
			} else if strings.EqualFold(param.valor, "WF") {
				fit1 = 'W'
			} else {
				fmt.Println("ERROR: El parametro fit no admite este valor -> " + param.valor)
				error = true
				break
			}
		} else if strings.EqualFold(param.name, "DELETE") {
			if strings.EqualFold(param.valor, "FAST") {
				tdelete = param.valor
				bandelete = true
			} else if strings.EqualFold(param.valor, "FULL") {
				tdelete = param.valor
				bandelete = true
			} else {
				fmt.Println("ERROR: El parametro delete no admite este valor -> " + param.valor)
				error = true
				break
			}
		} else if strings.EqualFold(param.name, "name") {
			name = param.valor
			bandname = true
		} else if strings.EqualFold(param.name, "ADD") {
			addP, _ := strconv.ParseInt(param.valor, 10, 64)
			add1 = addP
			bandadd = true
		} else {
			fmt.Println("ERROR: El comando fdisk no admite este parametro -> " + param.name)
			error = true
			break
		}
	}

	if !error {
		if bandpath {
			if bandname {
				if bandsize {
					if bandadd || bandelete {
						fmt.Println("Parametro Delete o ADD no soportado con size definido")
					} else {
						if bandtype {
							if type1 == 'P' {
								crearParticionP(path, name, size1, fit1, unit1)
							} else if type1 == 'E' {
								crearParticionE(path, name, size1, fit1, unit1)
							} else if type1 == 'L' {
								crearParticonL(path, name, size1, fit1, unit1)
							}

						} else {
							//path string, name string, size int64, fit byte, unit byte
							crearParticionP(path, name, size1, fit1, unit1)
						}
					}

				} else if bandadd {
					if bandsize || bandelete {
						fmt.Println("ERROR: Parametro Delete o Size no soportado con ADD definido")
					} else {
						agregarQuitarParticion(path, name, add1, unit1)
					}

				} else if bandelete {
					if bandsize || bandadd {
						fmt.Println("ERROR: Parametro ADD o Size no soportado con Delete definido")
					} else {
						eliminarParticion(path, name, tdelete)
					}
				}

			} else {
				fmt.Println("ERROR: El parametro name no esta definido")
			}

		} else {
			fmt.Println("ERROR: El parametro path no esta definido")
		}
	}

}

func ejecutarMOUNT(lisatributo []atributo) {
	var bandname bool = false
	var bandpaht bool = false
	var error bool = false
	var path string = ""
	var name string = ""
	for _, param := range lisatributo {
		if strings.EqualFold(param.name, "path") {
			bandpaht = true
			path = param.valor
		} else if strings.EqualFold(param.name, "name") {
			bandname = true
			name = param.valor
		} else {
			error = true
			fmt.Println("ERROR: El comando mount no admite este parametro -> " + param.name)
		}
	}

	if !error {
		if len(lisatributo) > 0 {
			if bandpaht {
				if bandname {
					//fmt.Println("estoy hasta la llamda")
					mount(path, name)
				} else {
					fmt.Println("ERROR: El parametro name no esta definido")
				}
			} else {
				fmt.Println("ERROR: El parametro path no esta definido")
			}
		} else {
			//fmt.Println("entre en mostrar")
			mostrarMon()
		}
	}

}

func ejecutarUNMOUNT(lisatributo []atributo) {
	for _, param := range lisatributo {
		if param.name[0] == 'i' || param.name[0] == 'I' {
			if param.name[1] == 'd' || param.name[1] == 'D' {
				_, ok := listamontada[param.valor]
				if ok {
					delete(listamontada, param.valor)
					fmt.Println("Particon desmontada con exito")
				} else {
					fmt.Println("ERROR: La particion no se encuetra montada")
				}
			} else {
				fmt.Println("ERROR: El comando unmount no admite el parametro -> " + param.name)
				break
			}
		} else {
			fmt.Println("ERROR: El comando unmount no admite el parametro -> " + param.name)
			break
		}
	}
}

func ejecutarREP(lisatributo []atributo) {
	var path string = ""
	var name string = ""
	var id string = ""
	var ruta string = ""
	var error bool = false
	var banpath bool = false
	var banname bool = false
	var banid bool = false

	for _, param := range lisatributo {
		if strings.EqualFold(param.name, "path") {
			banpath = true
			path = param.valor
		} else if strings.EqualFold(param.name, "name") {
			banname = true
			name = param.valor
		} else if strings.EqualFold(param.name, "id") {
			banid = true
			id = param.valor
		} else if strings.EqualFold(param.name, "ruta") {
			ruta = param.valor
		} else {
			error = true
			fmt.Println("ERROR El comando rep no admite el parametro -> " + param.name)
			break
		}
	}

	if !error {
		if banpath {
			if banname {
				if banid {
					valor, ok := listamontada[id]
					if ok {
						find := strings.LastIndexByte(path, '/')
						carpeta := path[0:find]
						find1 := strings.LastIndexByte(path, '.')
						extend := path[find1+1 : len(path)]
						if strings.EqualFold(name, "disk") {
							reportedisco(valor.Path, carpeta, path, extend)
						} else if strings.EqualFold(name, "mbr") {
							reportembr(valor.Path, carpeta, path, extend)
						} else {
							fmt.Print(ruta)
							fmt.Println("ERROR: el nombre del reporte no esta defido o no existe")

						}
					} else {
						fmt.Println("ERROR: no exite el id en las particones montadas")
					}

				} else {
					fmt.Println("Parametro id no definido")
				}
			} else {
				fmt.Println("ERROR: parametro name no definido")
			}
		} else {
			fmt.Println("ERROR: parametro path no definido")
		}
	}
}

func rmdisk(path string) {
	fi, err := os.Open(path) // Para acceso de lectura.
	if err == nil {
		fi.Close()
		println("Esta seguro que desea elimiar el disco (y/n)?  ")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		eleccion := strings.TrimRight(entrada, "\r\n")
		if eleccion == "Y" || eleccion == "y" {
			err := os.Remove(path)
			if err != nil {
				fmt.Println("ERROR: fatal error al tratar de eliminar el disco")
				fmt.Println(err)
			} else {
				fmt.Println("El disco fue eliminado con exito")
			}
		} else {
			fmt.Println("Eliminacion de disco abortada. ")
		}

	} else {
		fi.Close()
		fmt.Println("ERROR: No se ha encontrado el disco en la ruta especificada. (Disco no existe o ruta incorrecta)")
		fmt.Println(err)
	}
}

func crearParticionP(path string, name string, size int64, fit byte, unit byte) {
	var fitp byte
	var sizep int64
	var pathT string = path
	//verificamos que tipo de fit es el que vamos a aplicar
	if fit != 'N' {
		fitp = fit
	} else {
		fitp = 'W'
	}
	// verificamos si vino el parametro unit y definimos el tamaño
	if unit != 'N' {
		if unit == 'b' || unit == 'B' {
			sizep = size
		} else if unit == 'k' || unit == 'K' {
			sizep = size * 1024
		} else if unit == 'm' || unit == 'M' {
			sizep = size * 1024 * 1024
		}
	} else {
		sizep = size * 1024
	}
	if sizep <= 0 {
		fmt.Println(sizep)
		fmt.Println("ERROR:el parametro size solo acepta valores mayores a cero")
		return
	}

	mbr := mBR{}
	//fmt.Println(pathT)
	file, err := os.OpenFile(pathT, os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR: Disco no encontrado en la ruta especificada")
	} else {
		var dispParticion bool = false
		var numParticion int64 = 0

		file.Seek(0, 0)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)
		//fmt.Println(mbr.Mbrtamano)
		//fmt.Println(bytesToString(mbr.Mbrfechacreacion[:]))
		for i := 0; i < 4; i++ {
			//fmt.Println("entre a buscar el n0 de particon")
			if mbr.Particiones[i].Partstart == -1 || (mbr.Particiones[i].Partstatus == 'N' && mbr.Particiones[i].Partsize >= sizep) {
				dispParticion = true
				numParticion = int64(i)
				//fmt.Println(numParticion)
				break
			}
		}

		if dispParticion {
			var usingspace int64 = 0
			for j := 0; j < 4; j++ {
				if mbr.Particiones[j].Partstatus == 'S' || mbr.Particiones[j].Partstatus == 'M' {
					usingspace += mbr.Particiones[j].Partsize
				}
			}

			if (mbr.Mbrtamano - usingspace) >= sizep {
				//fmt.Println("entre al espacio")
				if !existeparticion(path, name) {
					//fmt.Println("pase de buscar name")
					//debido a lo escrito que solo seria el primer ajuste
					if numParticion == 0 {
						mbr.Particiones[numParticion].Partstart = int64(binary.Size(mbr))
					} else {
						mbr.Particiones[numParticion].Partstart = mbr.Particiones[numParticion-1].Partstart + mbr.Particiones[numParticion-1].Partsize
					}

					mbr.Particiones[numParticion].Parttype = 'P'
					mbr.Particiones[numParticion].Partfit = fitp
					mbr.Particiones[numParticion].Partsize = sizep
					mbr.Particiones[numParticion].Partstatus = 'S'
					mbr.Particiones[numParticion].Formateada = -1
					copy(mbr.Particiones[numParticion].Partname[:], name)

					//procedemos a guardar los comabios en el mbr
					file.Seek(0, 0)
					s := &mbr
					var binario2 bytes.Buffer
					binary.Write(&binario2, binary.BigEndian, s)
					escribirBytes(file, binario2.Bytes())
					//fin guardar combios en mbr
					fmt.Println("Particion Primaria Creada con exito ")

				} else {
					fmt.Println("ERROR: ya existe una particion con el mismo nombre")
				}

			} else {
				fmt.Println("ERROR: Espacio insuficiente para crear la particon Primaria")
			}

		} else {
			fmt.Println("ERROR: El numero de particones primarias llego a su limite")
		}

	}

}
func leerBytes(file *os.File, number int) []byte {
	//fmt.Println(number)
	bytes := make([]byte, number) //array de bytes

	_, err := file.Read(bytes) // Leido -> bytes
	if err != nil {
		fmt.Println("es la mierda de la lectura")
		log.Fatal(err)
	}

	return bytes
}
func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}

func bytesToString(data []byte) string {
	return string(data[:])
}
func existeparticion(path string, name string) bool {
	var pathT string = path
	var extendida int = -1
	file, err := os.OpenFile(pathT, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR: Disco no encontrado en la ruta especificada")
	} else {
		mbr := mBR{}
		file.Seek(0, 0)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)
		var nametem [16]byte
		for i := 0; i < 4; i++ {
			//fmt.Println("buscando nombre")
			//fmt.Println(bytesToString(mbr.Particiones[i].Partname[:]))
			//fmt.Println(mbr.Particiones[i].Partname)
			//fmt.Println(name)
			copy(nametem[:], name)
			//bytes.Compare(nametem[:], mbr.Particiones[i].Partname[:]) == 0
			if strings.EqualFold(bytesToString(nametem[:]), bytesToString(mbr.Particiones[i].Partname[:])) {
				return true
			} else if mbr.Particiones[i].Parttype == 'E' {
				extendida = i
			}
		}

		if extendida > -1 {
			//fmt.Println("entre a la extendida saber por que")
			file.Seek(mbr.Particiones[extendida].Partstart, 0)
			ebr := eBR{}
			for {
				var numb3 int = binary.Size(ebr)
				data1 := leerBytes(file, numb3)
				buffer1 := bytes.NewBuffer(data1)
				err = binary.Read(buffer1, binary.BigEndian, &ebr)
				pos, err1 := file.Seek(0, os.SEEK_CUR)
				if err != nil || err1 != nil || pos >= (mbr.Particiones[extendida].Partsize+mbr.Particiones[extendida].Partstart) {
					break
				}
				if strings.EqualFold(bytesToString(ebr.Partname[:]), bytesToString(nametem[:])) {
					return true
				}
				if ebr.Partnext == -1 {
					break
				} else {
					file.Seek(ebr.Partnext, os.SEEK_SET)
				}
			}

		}
		//fmt.Println("estoy en el reotono")
		return false

	}
	return false
}

func crearParticionE(path string, name string, size int64, fit byte, unit byte) {
	var fitp byte
	var sizep int64
	//verificamos que tipo de fit es el que vamos a aplicar
	if fit != 'N' {
		fitp = fit
	} else {
		fitp = 'W'
	}
	// verificamos si vino el parametro unit y definimos el tamaño
	if unit != 'N' {
		if unit == 'b' || unit == 'B' {
			sizep = size
		} else if unit == 'k' || unit == 'K' {
			sizep = size * 1024
		} else if unit == 'm' || unit == 'M' {
			sizep = size * 1024 * 1024
		}
	} else {
		sizep = size * 1024
	}
	if sizep <= 0 {
		fmt.Println(sizep)
		fmt.Println("ERROR:el parametro size solo acepta valores mayores a cero")
		return
	}

	mbr := mBR{}
	//fmt.Println(pathT)
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR: Disco no encontrado en la ruta especificada")
	} else {
		var dispParticion bool = false
		var numParticion int64 = 0
		var usginspace int64 = 0
		//para la lectura del mbr
		file.Seek(0, 0)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)

		for i := 0; i < 4; i++ {
			if mbr.Particiones[i].Parttype == 'E' {
				fmt.Println("ERROR: Ya existe una particion Ectendida solo se puede tener una en un disco")
				return
			}
			if !dispParticion {
				if mbr.Particiones[i].Partstart == -1 || (mbr.Particiones[i].Partstatus == 'N' && mbr.Particiones[i].Partsize >= sizep) {
					numParticion = int64(i)
					dispParticion = true
				}
			}

			if mbr.Particiones[i].Partstatus == 'S' || mbr.Particiones[i].Partstatus == 'M' {
				usginspace += mbr.Particiones[i].Partsize
			}
		}

		if dispParticion {
			if (mbr.Mbrtamano - usginspace) >= sizep {
				if !existeparticion(path, name) {
					if numParticion == 0 {
						mbr.Particiones[numParticion].Partstart = int64(binary.Size(mbr))
					} else {
						mbr.Particiones[numParticion].Partstart = mbr.Particiones[numParticion-1].Partstart + mbr.Particiones[numParticion-1].Partsize

					}

					mbr.Particiones[numParticion].Parttype = 'E'
					mbr.Particiones[numParticion].Partfit = fitp
					mbr.Particiones[numParticion].Partsize = sizep
					mbr.Particiones[numParticion].Partstatus = 'S'
					mbr.Particiones[numParticion].Formateada = -1
					copy(mbr.Particiones[numParticion].Partname[:], name)
					//procedemos a guardar los comabios en el mbr
					file.Seek(0, 0)
					s := &mbr
					var binario2 bytes.Buffer
					binary.Write(&binario2, binary.BigEndian, s)
					escribirBytes(file, binario2.Bytes())
					//fin guardar combios en mbr

					//inicio de guardar el ebr para la primera logica
					file.Seek(mbr.Particiones[numParticion].Partstart, os.SEEK_SET)

					ebr := eBR{}
					ebr.Partfit = fitp
					ebr.Partstatus = 'S'
					ebr.Partstart = mbr.Particiones[numParticion].Partstart
					ebr.Partsize = -1
					ebr.Partnext = -1
					copy(ebr.Partname[:], "")

					seb := &ebr
					var binarioeb bytes.Buffer
					binary.Write(&binarioeb, binary.BigEndian, seb)
					escribirBytes(file, binarioeb.Bytes())
					//fin de guardar ebr

					fmt.Println("Particon Extendida creada con exito")

				} else {
					fmt.Println("ERROR: Ya existe una particion con el mismo nombre")
				}

			} else {
				fmt.Println("ERROR: No es posible crear la particion exede el espacio disponible")
			}
		} else {
			fmt.Println("ERROR: Ya existen 4 pariciones, No es posible crear otra del tipo Extendida")
		}

	}
}

func crearParticonL(path string, name string, size int64, fit byte, unit byte) {
	var fitp byte
	var sizep int64
	//verificamos que tipo de fit es el que vamos a aplicar
	if fit != 'N' {
		fitp = fit
	} else {
		fitp = 'W'
	}
	// verificamos si vino el parametro unit y definimos el tamaño
	if unit != 'N' {
		if unit == 'b' || unit == 'B' {
			sizep = size
		} else if unit == 'k' || unit == 'K' {
			sizep = size * 1024
		} else if unit == 'm' || unit == 'M' {
			sizep = size * 1024 * 1024
		}
	} else {
		sizep = size * 1024
	}
	if sizep <= 0 {
		fmt.Println(sizep)
		fmt.Println("ERROR:el parametro size solo acepta valores mayores a cero")
		return
	}

	mbr := mBR{}
	//fmt.Println(pathT)
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR: Disco no encontrado en la ruta especificada")
	} else {
		var numextend int = -1
		//para la lectura del mbr
		file.Seek(0, 0)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)

		for i := 0; i < 4; i++ {
			if mbr.Particiones[i].Parttype == 'E' {
				numextend = i
				break
			}
		}
		if existeparticion(path, name) {
			fmt.Println("ERROR: ya existe una particion con el mismo nombre")

		} else {
			if numextend != -1 {
				var comienzo int64 = mbr.Particiones[numextend].Partstart
				ebr := eBR{}
				file.Seek(comienzo, os.SEEK_SET)
				datae := leerBytes(file, binary.Size(ebr))
				buffere := bytes.NewBuffer(datae)
				err = binary.Read(buffere, binary.BigEndian, &ebr)
				if ebr.Partsize == -1 {
					//fmt.Println("primera logica que se va crear")
					//debe ser la primera particon logica
					if mbr.Particiones[numextend].Partsize > sizep {
						ebr.Partstatus = 'S'
						posact, _ := file.Seek(0, os.SEEK_CUR)
						var ppstar int64 = int64(posact) - int64(binary.Size(ebr))
						ebr.Partstart = ppstar
						ebr.Partfit = fitp
						ebr.Partsize = sizep
						ebr.Partnext = -1
						copy(ebr.Partname[:], name)
						file.Seek(mbr.Particiones[numextend].Partstart, os.SEEK_SET)
						seb := &ebr
						var binarioeb bytes.Buffer
						binary.Write(&binarioeb, binary.BigEndian, seb)
						escribirBytes(file, binarioeb.Bytes())
						fmt.Println("Particion Logica Creada con Exito")

					} else {
						fmt.Println("ERROR: el tamaño de la particon logica exede el espacio disponible")
					}
				} else {
					for {
						pos, _ := file.Seek(0, os.SEEK_CUR)
						//fmt.Print("posicion del file -> ")
						//fmt.Println(pos)
						if (ebr.Partnext == -1) || pos > (mbr.Particiones[numextend].Partsize+mbr.Particiones[numextend].Partstart) {
							break
						} else {
							file.Seek(ebr.Partnext, os.SEEK_SET)
							//fmt.Println("buacando el siguiente")
							var numb3 int = binary.Size(ebr)
							data1 := leerBytes(file, numb3)
							buffer1 := bytes.NewBuffer(data1)
							err = binary.Read(buffer1, binary.BigEndian, &ebr)
						}

					}

					var suficiente int64 = ebr.Partstart + ebr.Partsize + sizep
					if suficiente < mbr.Particiones[numextend].Partstart+mbr.Particiones[numextend].Partsize {
						ebr.Partnext = ebr.Partsize + ebr.Partstart
						pos11, _ := file.Seek(0, os.SEEK_CUR)
						var poss1 int64 = pos11 - int64(binary.Size(ebr))
						//fmt.Println(poss1)
						file.Seek(poss1, os.SEEK_SET)
						seb1 := &ebr
						var binarioeb1 bytes.Buffer
						binary.Write(&binarioeb1, binary.BigEndian, seb1)
						escribirBytes(file, binarioeb1.Bytes())

						file.Seek(ebr.Partsize+ebr.Partstart, os.SEEK_SET)

						copy(ebr.Partname[:], name)
						ebr.Partstatus = 'S'
						ebr.Partfit = fitp
						posfi, _ := file.Seek(0, os.SEEK_CUR)
						ebr.Partstart = int64(posfi)
						ebr.Partnext = -1
						ebr.Partsize = sizep

						sebfi := &ebr
						var binariofi bytes.Buffer
						binary.Write(&binariofi, binary.BigEndian, sebfi)
						escribirBytes(file, binariofi.Bytes())

						fmt.Println("Particion Logica Creada con Exito")

					} else {
						fmt.Println("El tamaño de la particon logica exede el espacio disponible")
					}

				}

			} else {
				fmt.Println("ERROR: Es necesario una paticion extendida para crear particones logicas ")
			}
		}
	}
}

func agregarQuitarParticion(path string, name string, add int64, unit byte) {

}

func eliminarParticion(path string, name string, forma string) {
	var indice int64 = -1
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR: No se encuetra el disco espefificado")
		log.Fatal(err)
	} else {
		var montada = false
		if !montada {
			indice = buscarParticonPE(path, name)
			if indice > -1 { // fue encontrada en las particones principales
				println("Esta seguro que desea elimiar la particion del  disco (y/n)?  ")
				reader := bufio.NewReader(os.Stdin)
				entrada, _ := reader.ReadString('\n')
				eleccion := strings.TrimRight(entrada, "\r\n")
				if eleccion == "Y" || eleccion == "y" {
					mbr := mBR{}
					file.Seek(0, 0)
					var numb2 int = binary.Size(mbr)
					data := leerBytes(file, numb2)
					buffer := bytes.NewBuffer(data)
					err = binary.Read(buffer, binary.BigEndian, &mbr)
					var auxty byte = mbr.Particiones[indice].Parttype
					if strings.EqualFold(forma, "FAST") {
						mbr.Particiones[indice].Partstatus = 'd'
						copy(mbr.Particiones[indice].Partname[:], "")
						file.Seek(0, 0)
						s := &mbr
						var binario2 bytes.Buffer
						binary.Write(&binario2, binary.BigEndian, s)
						escribirBytes(file, binario2.Bytes())
						if auxty == 'P' {
							fmt.Println("Particion Primaria Eliminada con Exito")
						} else {
							fmt.Println("Particion Extendida Eliminada con exito")
						}

					} else {
						mbr.Particiones[indice].Partstatus = 'd'
						copy(mbr.Particiones[indice].Partname[:], "")
						file.Seek(0, 0)
						s := &mbr
						var binario2 bytes.Buffer
						binary.Write(&binario2, binary.BigEndian, s)
						escribirBytes(file, binario2.Bytes())

						file.Seek(mbr.Particiones[indice].Partstart, os.SEEK_SET)
						var otro int8 = 0
						var rebina bytes.Buffer
						binary.Write(&rebina, binary.BigEndian, otro)
						escribirBytes(file, rebina.Bytes())
						rebina.Reset()
						file.Seek(mbr.Particiones[indice].Partstart+2, os.SEEK_SET)
						binary.Write(&rebina, binary.BigEndian, otro)
						escribirBytes(file, rebina.Bytes())
						if auxty == 'P' {
							fmt.Println("Particion Primaria Eliminada con Exito")
						} else {
							fmt.Println("Particion Extendida Eliminada con exito")
						}
					}

				} else {
					fmt.Println("Operacion de Eliminar Particion fue Abortada")
				}

			} else {
				//hay que buscar en las logicas
			}

		} else {
			fmt.Println("ERROR: No es posible eliminar la particion lla que se encuetra montada")
			fmt.Println("Primero tendra que desmontar la particion")
		}
	}
	buscarParticonPE(path, name)
}
func buscarParticonPE(path string, name string) int64 {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil { //validar que no sea nulo.
		fmt.Println("ERROR:no se encuentra disco especificado")
		log.Fatal(err)
	} else {
		mbr := mBR{}
		file.Seek(0, 0)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)
		var temname [16]byte
		copy(temname[:], name)
		for i := 0; i < 4; i++ {
			if mbr.Particiones[i].Partstatus != 'd' {
				//fmt.Println(string(mbr.Particiones[i].Partname[:]))
				if strings.EqualFold(bytesToString(mbr.Particiones[i].Partname[:]), bytesToString(temname[:])) {
					return int64(i)
				}
			}
		}
	}
	return -1
}

func buscarParticionL(path string, name string) int64 {
	var indice int64 = -1
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil { //validar que no sea nulo.
		log.Fatal(err)
	} else {
		mbr := mBR{}
		file.Seek(0, 0)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)
		for i := 0; i < 4; i++ {
			if mbr.Particiones[i].Parttype == 'E' {
				indice = int64(i)
				break
			}
		}

		if indice > -1 {
			ebr := eBR{}
			file.Seek(mbr.Particiones[indice].Partstart, os.SEEK_SET)
			datae := leerBytes(file, binary.Size(ebr))
			buffere := bytes.NewBuffer(datae)
			err = binary.Read(buffere, binary.BigEndian, &ebr)
			//fmt.Println(ebr.Partname[:])
			//fmt.Println(bytesToString(ebr.Partname[:]))
			if ebr.Partsize == -1 {
				return -1
			}
			for {
				pos, _ := file.Seek(0, os.SEEK_CUR)
				var temname [16]byte
				copy(temname[:], name)
				if strings.EqualFold(bytesToString(ebr.Partname[:]), bytesToString(temname[:])) {
					delt, _ := file.Seek(0, os.SEEK_CUR)
					indice = delt - int64(binary.Size(mbr))
					break
				}
				if err != nil || ebr.Partnext == -1 || pos > (mbr.Particiones[indice].Partsize+mbr.Particiones[indice].Partstart) {
					indice = -1
					break
				} else {

					file.Seek(ebr.Partnext, os.SEEK_SET)
					var numb3 int = binary.Size(ebr)
					data1 := leerBytes(file, numb3)
					buffer1 := bytes.NewBuffer(data1)
					err = binary.Read(buffer1, binary.BigEndian, &ebr)
				}
			}

		}

	}
	return indice
}

func mount(path string, name string) {
	//fmt.Println(path)
	//fmt.Println(name)
	var indicePar int64 = buscarParticonPE(path, name)
	if indicePar > -1 {
		file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
		defer file.Close()
		if err != nil { //validar que no sea nulo.
			fmt.Println("ERROR: No se encuetra el disco especificado")
			log.Fatal(err)
		} else {
			mbr := mBR{}
			file.Seek(0, 0)
			var numb2 int = binary.Size(mbr)
			data := leerBytes(file, numb2)
			buffer := bytes.NewBuffer(data)
			err = binary.Read(buffer, binary.BigEndian, &mbr)

			mbr.Particiones[indicePar].Partstatus = 'M'

			file.Seek(0, 0)
			s := &mbr
			var binario2 bytes.Buffer
			binary.Write(&binario2, binary.BigEndian, s)
			escribirBytes(file, binario2.Bytes())
			//fin guardar combios en mbr

			var letra byte = buscarL(path, name)

			if letra == '1' {
				fmt.Println("ERROR: la particcion ya esta montada")
			} else {
				var numero int64 = buscarN(path)
				var id string = "vd" + string(letra) + strconv.Itoa(int(numero))
				nodo1 := nodom{Path: path, Name: name, Numero: numero, Letra: letra}
				listamontada[id] = nodo1
				fmt.Println("Particon Montada con Exito")
				//fmt.Println(listamontada)

			}

		}

	} else {
		indicePar = buscarParticionL(path, name)
		if indicePar != -1 {
			file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
			defer file.Close()
			if err != nil { //validar que no sea nulo.
				fmt.Println("ERROR: No se encuetra el disco especificado")
				log.Fatal(err)
			} else {
				ebr := eBR{}
				file.Seek(indicePar, os.SEEK_SET)
				var numb2 int = binary.Size(ebr)
				data := leerBytes(file, numb2)
				buffer := bytes.NewBuffer(data)
				err = binary.Read(buffer, binary.BigEndian, &ebr)

				ebr.Partstatus = 'M'

				file.Seek(indicePar, os.SEEK_SET)
				s := &ebr
				var binario2 bytes.Buffer
				binary.Write(&binario2, binary.BigEndian, s)
				escribirBytes(file, binario2.Bytes())
				//fin guardar combios en mbr

				var letra byte = buscarL(path, name)

				if letra == '1' {
					fmt.Println("ERROR: la particcion ya esta montada")
				} else {
					var numero int64 = buscarN(path)
					var id string = "vd" + string(letra) + strconv.Itoa(int(numero))
					nodo1 := nodom{Path: path, Name: name, Numero: numero, Letra: letra}
					listamontada[id] = nodo1
					fmt.Println("Particon Logica Montada con Exito")
					//fmt.Println(listamontada)

				}

			}
		} else {
			fmt.Println("ERROR: No se encuetra la particon a montar")
		}
	}
}

func buscarL(path string, name string) byte {
	var response byte = 'a'
	for _, valor := range listamontada {
		if valor.Path == path && strings.EqualFold(valor.Name, name) {
			response = '1'
			break
		} else {
			if valor.Path == path {
				return valor.Letra
			} else if response <= valor.Letra {
				response++
			}
		}
	}
	return response
}
func buscarN(path string) int64 {
	var response int64 = 1
	for _, valor := range listamontada {
		if valor.Path == path {
			response++
		}
	}
	return response
}

func mostrarMon() {
	fmt.Println(":::::::::::::::::::::::::::::::::")
	fmt.Println(":       Particiones Montadas    :")
	fmt.Println("---------------------------------")
	fmt.Println(":      Nombre    |    ID        :")
	fmt.Println(":::::::::::::::::::::::::::::::::")
	for key, valor := range listamontada {
		fmt.Println(":   " + valor.Name + "          " + key)
		fmt.Println("---------------------------------")
	}
}

func reportedisco(path string, carpeta string, salida string, extend string) {

}

func reportembr(path string, carpeta string, salida string, extend string) {
	crearDirectorioF(carpeta)
	mbr := mBR{}
	ebr := eBR{}
	var numextend int = -1
	var auxdot string = ""
	//fmt.Println(pathT)
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR: Al crear el Reporte MBR disco no encontrado")
	} else {
		auxdot = "digraph G{ \n"
		auxdot += "subgraph cluster{\n label=\"MBR\" \n"
		auxdot += "rmbr[shape=box,label=<\n"
		auxdot += "<table border='0' cellborder='1' cellspacing='0' width='300'  height='200' >\n"
		auxdot += "<tr>  <td width='150'> <b>Nombre</b> </td> <td width='150'> <b>Valor</b> </td>  </tr>\n"

		file.Seek(0, os.SEEK_SET)
		var numb2 int = binary.Size(mbr)
		data := leerBytes(file, numb2)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &mbr)
		auxdot += "<tr>  <td><b>mbr_tamaño</b></td><td>"
		auxdot += strconv.Itoa(int(mbr.Mbrtamano)) + "</td>  </tr>\n"
		auxdot += "<tr>  <td><b>mbr_fecha_creacion</b></td> <td>" + string(mbr.Mbrfechacreacion[:])
		auxdot += "</td>  </tr>\n"
		auxdot += "<tr>  <td><b>mbr_disk_signature</b></td> <td>" + strconv.Itoa(int(mbr.Mbrdisksignature)) + "</td>  </tr>\n"
		//auxdot += "<tr>  <td><b>Disk_fit</b></td> <td>" + string(mbr.) + "</td>  </tr>\n"

		for i := 0; i < 4; i++ {
			if mbr.Particiones[i].Partstart != -1 && mbr.Particiones[i].Partstatus != 'N' {
				if mbr.Particiones[i].Parttype == 'E' {
					numextend = i
				}

				auxdot += "<tr>  <td><b>part_status_" + strconv.Itoa(i+1)
				auxdot += "</b></td> <td>" + string(mbr.Particiones[i].Partstatus)
				auxdot += "</td>  </tr>\n"
				auxdot += "<tr>  <td><b>part_type_" + strconv.Itoa(i+1) + "</b></td> <td>" + string(mbr.Particiones[i].Parttype) + "</td>  </tr>\n"
				auxdot += "<tr>  <td><b>part_fit_" + strconv.Itoa(i+1) + "</b></td> <td>" + string(mbr.Particiones[i].Partfit) + "</td>  </tr>\n"
				auxdot += "<tr>  <td><b>part_start_" + strconv.Itoa(i+1) + "</b></td> <td>" + strconv.Itoa(int(mbr.Particiones[i].Partstart)) + "</td>  </tr>\n"
				auxdot += "<tr>  <td><b>part_size_" + strconv.Itoa(i+1) + "</b></td> <td>" + strconv.Itoa(int(mbr.Particiones[i].Partsize)) + "</td>  </tr>\n"
				n := bytes.IndexByte(mbr.Particiones[i].Partname[:], 0)
				auxdot += "<tr>  <td><b>part_name_" + strconv.Itoa(i+1) + "</b></td> <td>" + bytesToString(mbr.Particiones[i].Partname[:n]) + "</td>  </tr>\n"

			}
		}
		auxdot += "</table>\n >];\n}\n"

		if numextend > -1 {
			var indebr int = 1
			file.Seek(mbr.Particiones[numextend].Partstart, os.SEEK_SET)
			for {
				var numb3 int = binary.Size(ebr)
				data1 := leerBytes(file, numb3)
				buffer1 := bytes.NewBuffer(data1)
				err = binary.Read(buffer1, binary.BigEndian, &ebr)
				pos, err1 := file.Seek(0, os.SEEK_CUR)
				if err != nil || err1 != nil || pos >= (mbr.Particiones[numextend].Partsize+mbr.Particiones[numextend].Partstart) {
					break
				}
				//comienzo de graphivz
				if ebr.Partstatus != 'N' {
					auxdot += "subgraph cluster_" + strconv.Itoa(indebr) + "{\n label=\"EBR_" + strconv.Itoa(indebr) + "\"\n"
					auxdot += "nebr_" + strconv.Itoa(indebr) + "[shape=box, label=<\n"
					auxdot += "<table border='0' cellborder='1' cellspacing='0'  width='300' height='160' >\n "
					auxdot += "<tr>  <td width='150'><b>Nombre</b></td> <td width='150'><b>Valor</b></td>  </tr>\n"

					//--------------------------------------------------------------------------------------------------------
					auxdot += "<tr>  <td><b>part_status_L"
					auxdot += "</b></td> <td>" + string(ebr.Partstatus)
					auxdot += "</td>  </tr>\n"
					auxdot += "<tr>  <td><b>part_next_L </b></td> <td>" + strconv.Itoa(int(ebr.Partnext)) + "</td>  </tr>\n"
					auxdot += "<tr>  <td><b>part_fit_L</b></td> <td>" + string(ebr.Partfit) + "</td>  </tr>\n"
					auxdot += "<tr>  <td><b>part_start_L</b></td> <td>" + strconv.Itoa(int(ebr.Partstart)) + "</td>  </tr>\n"
					auxdot += "<tr>  <td><b>part_size_L </b></td> <td>" + strconv.Itoa(int(ebr.Partsize)) + "</td>  </tr>\n"
					n1 := bytes.IndexByte(ebr.Partname[:], 0)
					auxdot += "<tr>  <td><b>part_name_L</b></td> <td>" + bytesToString(ebr.Partname[:n1]) + "</td>  </tr>\n"
					auxdot += "</table>\n >];\n}\n"
					//---------------------------------------------------------------------------------------------------
					indebr++

				}
				// fin graphivz
				if ebr.Partnext == -1 {
					break
				} else {
					file.Seek(ebr.Partnext, os.SEEK_SET)
				}
			}

		}

		auxdot += "}\n"

		dotA, _ := os.OpenFile("MBR1.dot", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer dotA.Close()

		_, err := dotA.WriteString(auxdot)
		if err != nil {
			panic(err)
		}
		//comando1 := "dot -T" + extend + " MBR1.dot -o " + salida
		arg1 := "-T" + extend
		cmd := exec.Command("dot", arg1, "MBR1.dot", "-o", salida) // no need to call Output method here
		err1 := cmd.Run()
		if err1 != nil {
			fmt.Println("Entra por algun erro del cmd")
			log.Fatal(err)
		}
		fmt.Println("REPORTE MBR generado con exito")
		fmt.Println(extend)
		fmt.Println(salida)
	}
}
