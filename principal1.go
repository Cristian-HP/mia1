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
	mbrtamano        int64
	mbrfechacreacion time.Time
	mbrdisksignature int64
	particiones      [4]pARTICION
}

type pARTICION struct {
	partstatus byte
	parttype   byte
	partfit    byte
	partstart  int64
	partsize   int64
	partname   [16]byte
	formateada int64
}

type eBR struct {
	partstatus byte
	partfit    byte
	partstart  int64
	partsize   int64
	partnext   int64
	partname   [16]byte
}

//este apartado es para las variable globales para teneer por como siem
var listoken []token
var listcomand []comando
var tipo int64

//sirve para la mierda de las carpetas

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
	file, err := os.Create(aux1)
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
		var i int64
		d2 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i = 0; i < (size / 1024); i++ {
			//fmt.Fprintf(file, "%c", buffer1)
			file.Write(d2)
		}

		// Crear el MBR
		file.Seek(0, 0)

		mbr := mBR{}
		//memset(&mbr, 0, sizeof(MBR))
		mbr.mbrtamano = size
		mbr.mbrdisksignature = rand.Int63n(100)
		mbr.mbrfechacreacion = time.Now()
		//mbr.fit_disk = fitt
		for p := 0; p < 4; p++ {
			mbr.particiones[p].partstatus = 'N'
			mbr.particiones[p].partsize = 0
			mbr.particiones[p].parttype = 'N'
			mbr.particiones[p].partfit = 'N'
			mbr.particiones[p].partstart = -1
			mbr.particiones[p].formateada = -1
			copy(mbr.particiones[p].partname[:], "")
		}
		//fwrite(&mbr, sizeof(MBR), 1, file)
		s := &mbr
		var binario2 bytes.Buffer
		binary.Write(&binario2, binary.BigEndian, s)
		file.Write(binario2.Bytes())
		file.Close()
		file.Sync()
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
	fmt.Println(len(listcomand))
	for _, com := range listcomand {
		if strings.EqualFold(com.name, "exec") {
			fmt.Println("entre a ejecu")
			ejecutarEXEC(com.lisAtri)
		} else {
			//fmt.Println("esto es el de ejecutat")
			fmt.Println(com.name)
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
