package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Задача 1 -----------------------------------------------------------------------------
// Создайте программу, которая читает лог-файл server.log,
// подсчитывает количество строк, содержащих слово "error", и выводит это число.

func CountSubstrInFile(filename, substr string) (int, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("CountSubstrInFile: %w", err)
	}
	return strings.Count(string(bytes), substr), nil
}

func Ex1() {
	filename := "logfile.log"
	substr := "error"

	n, err := CountSubstrInFile(filename, substr)
	if err != nil {
		fmt.Println("Ex1:", err)
		return
	}
	fmt.Printf("File '%s' contains '%s' %d times\n", filename, substr, n)
}

// Задача 2 -----------------------------------------------------------------------------
// Напишите программу, которая принимает список имен файлов в текущей директории,
// объединяет их содержимое и сохраняет результат в новый файл combined.txt
// (только текстовые файлы)

func ForEachFile(dir string, callback func(file string) error) error {
	const fnName = "ForEachFile" // Got this from a standard library

	// Maybe os.ReadDir() already does that, not sure
	info, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("%s: %w", fnName, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s: '%s' is not a directory", fnName, dir)
	}

	// NOTE: filepath.WalkDir() can be used for recursive lookup
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("%s: %w", fnName, err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// filepath.Abs(entry.Name()) skips "newDirectory/", which is surprising
		// So we need to Join() with current directory name, and then use Abs()
		absPath, err := filepath.Abs(filepath.Join(dir, entry.Name()))
		if err != nil {
			return fmt.Errorf("%s: %w", fnName, err)
		}

		if err = callback(absPath); err != nil {
			return fmt.Errorf("%s: %w", fnName, err)
		}
	}
	return nil
}

func MergeAllTextFiles(dir string, outFile string) error {
	const fnName = "MergeAllTextFiles"

	var buffer []byte
	err := ForEachFile(dir, func(file string) error {
		if filepath.Ext(file) == ".txt" {
			moreBytes, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("%s: %w", "callback", err)
			}
			buffer = append(buffer, moreBytes...) // Very interesting use of "..."
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: %w", fnName, err)
	}
	return os.WriteFile(outFile, buffer, os.ModeAppend)
}

func Ex2() {
	err := MergeAllTextFiles("newDirectory", "combined.txt")
	if err != nil {
		fmt.Println("Ex2:", err)
		return
	}
}

func main() {
	fmt.Println("Ex1")
	Ex1()
	fmt.Println("\nEx2")
	Ex2()
}
