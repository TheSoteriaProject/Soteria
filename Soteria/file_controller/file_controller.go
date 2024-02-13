package file_controller

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// Other
	"Soteria/ignore_file_parser"
)

// Confirm File Controller Connection
func TestConnection() {
	fmt.Println("Testing File Controller Connection.")
}

// Show Slice Data for Debugging
func ShowSliceData(path []string) {
	for _, path := range path {
		fmt.Printf("%s\n", path)
	}
}

// Walk the Path Traversal
func WalkTheFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		file := GetAllFiles(path, info, err)
		files = append(files, file...)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
		return nil
	}

	return files
}

// Path Traversal
func GetAllFiles(path string, info os.FileInfo, err error) []string {
	files := []string{}

	if err != nil {
		fmt.Println("Error:", err)
		return files
	}

	if !info.IsDir() {
		files = append(files, path)
	}

	return files
}

func FilterFileExtensions(files []string, u_makefile bool, u_dockerfile bool, u_bash bool) []string {
	filtered_files := []string{}

	for _, files := range files {
		split := strings.Split(files, "/")
		extension := filepath.Ext(split[len(split)-1]) // MAY WANT FULL FILE PATH
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if u_makefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, files)
		}
		// Dockerfile Check
		if u_dockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, files)
		}
		// Bash Check
		if u_bash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, files)
		}
	}

	return filtered_files
}

func CompareFiles(files []string, ignored_files []string) []string {
	filtered_files := []string{} // Files that have passed filtering
	remove_files := []string{}

	// check if file is in ignore folder
	for _, file := range files {
		for _, i_file := range ignored_files {
			// Break up Token
			split_string := strings.SplitAfter(i_file, ":")
			front_of_split_string := split_string[0]
			end_of_split_string := split_string[1]
			if strings.TrimSpace(end_of_split_string) == "IgnoreFile" {
				// Trim the negative sign
				front_of_split_string = strings.TrimSpace(front_of_split_string)
				front_of_split_string = strings.TrimLeft(front_of_split_string, "-")
				front_of_split_string = strings.TrimRight(front_of_split_string, ":")

				// Genius or dum????
				/*
					trim_space_extra := strings.TrimSpace(front_of_split_string)
					dir_preface_path := trim_space_extra[:(len(trim_space_extra) - len(file))-1]
					remapped_name := dir_preface_path + "/" + file
				*/
				trim_space_extra := strings.TrimSpace(front_of_split_string) // Get filepath length of ignore
				file = strings.TrimSpace(file)

				range_var := -(len(trim_space_extra) - len(file))
				if range_var >= -1 {
					file_path := file[range_var:]
					// fmt.Println("Debug1: " + strings.TrimSpace(file) +":")
					// fmt.Println("Debug2: " + strings.TrimSpace(trim_space_extra) +":")
					// fmt.Println("Debug3: " + strings.TrimSpace(file_path) +":")

					if strings.TrimSpace(front_of_split_string) == strings.TrimSpace(file_path) {
						remove_files = append(remove_files, file_path) // May not be able to use file
					}
				}
			} else if strings.TrimSpace(end_of_split_string) == "IncludeFile" {
				// Trim the negative sign
				front_of_split_string = strings.TrimSpace(front_of_split_string)
				front_of_split_string = strings.TrimLeft(front_of_split_string, "+")
				front_of_split_string = strings.TrimRight(front_of_split_string, ":")

				// Genius or dum????
				trim_space_extra := strings.TrimSpace(front_of_split_string)
				// fmt.Println("Trim White Space: " + trim_space_extra)
				if (len(trim_space_extra) - len(file) - 1) > 0 {
					dir_preface_path := trim_space_extra[:(len(trim_space_extra)-len(file))-1]
					// fmt.Println("Parse Dir Path: " + dir_preface_path)
					remapped_name := dir_preface_path + "/" + file
					// fmt.Println("Remapped: " + remapped_name)
					// fmt.Println("Debug1: " + strings.TrimSpace(front_of_split_string) +":")
					// fmt.Println("Debug2: " + strings.TrimSpace(remapped_name) +":")
					if strings.TrimSpace(front_of_split_string) == strings.TrimSpace(remapped_name) {
						filtered_files = append(filtered_files, remapped_name)
					}
				}
			} else if strings.TrimSpace(end_of_split_string) == "IncludeFolder" {
				/*
				   // Trim the negative sign
				   front_of_split_string = strings.TrimSpace(front_of_split_string)
				   front_of_split_string = strings.TrimLeft(front_of_split_string, "+")
				   front_of_split_string = strings.TrimRight(front_of_split_string, ":")

				   // Genius or dum????
				   trim_space_extra := strings.TrimSpace(front_of_split_string)
				   // fmt.Println("Trim White Space: " + trim_space_extra)
				   if (len(trim_space_extra) - len(file)-1) > 0 {
				           dir_preface_path := trim_space_extra[:(len(trim_space_extra) - len(file))-1]
				           // fmt.Println("Parse Dir Path: " + dir_preface_path)
				           remapped_name := dir_preface_path + "/" + file
				           // fmt.Println("Remapped: " + remapped_name)
				           // fmt.Println("Debug1: " + strings.TrimSpace(front_of_split_string) +":")
				           // fmt.Println("Debug2: " + strings.TrimSpace(remapped_name) +":")
				           if strings.TrimSpace(front_of_split_string) == strings.TrimSpace(remapped_name) {
				                   filtered_files = append(filtered_files, front_of_split_string)
				           }
				   } */
				continue
			} else if strings.TrimSpace(end_of_split_string) == "IgnoreFolder" {
				/*
									// Trim the negative sign
					                                front_of_split_string = strings.TrimSpace(front_of_split_string)
					                                front_of_split_string = strings.TrimLeft(front_of_split_string, "-")
					                                front_of_split_string = strings.TrimRight(front_of_split_string, ":")

					                                // Genius or dum????
					                                trim_space_extra := strings.TrimSpace(front_of_split_string)
					                                // fmt.Println("Trim White Space: " + trim_space_extra)
					                                if (len(trim_space_extra) - len(file)-1) > 0 {
					                                        dir_preface_path := trim_space_extra[:(len(trim_space_extra) - len(file))-1]
					                                        // fmt.Println("Parse Dir Path: " + dir_preface_path)
					                                        remapped_name := dir_preface_path + "/" + file
					                                        // fmt.Println("Remapped: " + remapped_name)
					                                        // fmt.Println("Debug1: " + strings.TrimSpace(front_of_split_string) +":")
					                                        // fmt.Println("Debug2: " + strings.TrimSpace(remapped_name) +":")
					                                        if strings.TrimSpace(front_of_split_string) == strings.TrimSpace(remapped_name) {
					                                                filtered_files = append(filtered_files, front_of_split_string)
					                                        }
					                                }
				*/
				continue
			} else if strings.TrimSpace(end_of_split_string) == "IgnoreExtension" {
				/*
				   // Trim the negative sign
				   front_of_split_string = strings.TrimSpace(front_of_split_string)
				   front_of_split_string = strings.TrimLeft(front_of_split_string, "+")
				   front_of_split_string = strings.TrimRight(front_of_split_string, ":")

				   // Genius or dum????
				   trim_space_extra := strings.TrimSpace(front_of_split_string)
				   // fmt.Println("Trim White Space: " + trim_space_extra)
				   if (len(trim_space_extra) - len(file)-1) > 0 {
				           dir_preface_path := trim_space_extra[:(len(trim_space_extra) - len(file))-1]
				           // fmt.Println("Parse Dir Path: " + dir_preface_path)
				           remapped_name := dir_preface_path + "/" + file
				           // fmt.Println("Remapped: " + remapped_name)
				           // fmt.Println("Debug1: " + strings.TrimSpace(front_of_split_string) +":")
				           // fmt.Println("Debug2: " + strings.TrimSpace(remapped_name) +":")
				           if strings.TrimSpace(front_of_split_string) == strings.TrimSpace(remapped_name) {
				                   filtered_files = append(filtered_files, front_of_split_string)
				           }
				   } */
				continue
			} else {
				continue
			}
		}

		// issue is this file misses pre-face
		filtered_files = append(filtered_files, file)
	}
	// Check if file is ignored
	// Check if it needs to be appended because include in folder
	// Check if file is included so just add it
	// File not be in any case so just add it which is else case
	// With fron of new string maybe append and check???

	/*************************************/
	/************ BAD FIX ADJUST ********/
	/***********************************/
	var result_pool []string

	// Create a function to check if a file is in the remove_files slice
	isInRemoveFiles := func(file string) bool {
		for _, rf := range remove_files {
			f_range := len(file) - len(rf)
			if f_range > 0 { // May have to adjust
				// fmt.Println("Length", f_range)
				// fmt.Println("1:" + file[f_range:] + "|")
				// fmt.Println("2:" + rf + "|")
				file_adjust := file[f_range:]
				if file_adjust == rf {
					return true
				}
			}
		}
		return false
	}

	// Iterate through filtered_files and add non-removed files to result
	for _, file := range filtered_files {
		if !isInRemoveFiles(file) {
			result_pool = append(result_pool, file)
		}
	}
	fmt.Println("Remove Files")
	ShowSliceData(remove_files)
	fmt.Println("Result Pool")
	return result_pool
}

// Main Controller For File Controller
func FileController(path string) {
	// Get All Files / Walk The Directories
	files := WalkTheFiles(path)

	// Extract bash, Makefiles, and Dockerfiles
	// Set Variables u stands for use
	// Do in Main Controller
	u_bash := true
	u_dockerfile := true
	u_makefile := true

	extension_filtered_files := FilterFileExtensions(files, u_makefile, u_dockerfile, u_bash)
	ShowSliceData(extension_filtered_files)

	// Test Connection
	ignore_file_parser.TestConnection()

	// Ignore Cases w Tokens
	filter_cases := ignore_file_parser.FilterFiles()
	// ShowSliceData(filter_cases)

	// Compare Ignored and all files grabbed
	file_pool := CompareFiles(extension_filtered_files, filter_cases)
	ShowSliceData(file_pool)
}
