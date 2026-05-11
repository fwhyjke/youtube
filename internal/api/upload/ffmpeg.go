package upload

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

// func BuildFinalFile(bb BuildingBlocks, output string) {
// 	if bb.VPath != "" {
// 		if strings.Contains(bb.VCodec, "avc1") {
// 			if bb.APath != "" {
// 				ffmpegBuildMp4(bb.VPath, bb.APath, output)
// 			} else {
// 				ffmpegBuildMp4NoAudio(bb.VPath, output)
// 			}
// 		} else {
// 			if bb.APath != "" {
// 				ffmpegRecodeMp4(bb.VPath, bb.APath, output)
// 			} else {
// 				ffmpegRecodeMp4NoAudio(bb.VPath, output)
// 			}
// 		}
// 	} else if bb.APath != "" {
// 		ffmpegBuildMp3(bb.APath, output)
// 	} else {
// 		panic("Инвалид?")
// 	}
// }

// func ffmpegBuildMp4(vPath, aPath, outputFile string) {
// 	outputFile = fmt.Sprintf("data/%s.mp4", outputFile)
// 	cmd := exec.Command("ffmpeg", "-y",
// 		"-i", vPath,
// 		"-i", aPath,
// 		"-c", "copy",
// 		"-map", "0:v:0",
// 		"-map", "1:a:0",
// 		"-shortest",
// 		outputFile,
// 	)

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("\nОшибка при склейке: %v\n", err)
// 		return
// 	}

// 	os.Remove(vPath)
// 	os.Remove(aPath)

// 	fmt.Printf("\nГотово! Видео сохранено в: %s\n", outputFile)
// }

// func ffmpegRecodeMp4(vPath, aPath, outputFile string) {
// 	outputFile = fmt.Sprintf("data/%s.mp4", outputFile)
// 	cmd := exec.Command("ffmpeg", "-y",
// 		"-i", vPath,
// 		"-i", aPath,
// 		"-c:v", "libx264", // Перекодируем видео
// 		"-crf", "18", // Качество 18 - почти без потерь
// 		"-preset", "fast",
// 		"-c:a", "copy",
// 		"-map", "0:v:0",
// 		"-map", "1:a:0",
// 		"-shortest",
// 		outputFile,
// 	)

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("\nОшибка при склейке: %v\n", err)
// 		return
// 	}

// 	os.Remove(vPath)
// 	os.Remove(aPath)

// 	fmt.Printf("\nГотово! Видео сохранено в: %s\n", outputFile)
// }

// func ffmpegBuildMp4NoAudio(vPath, outputFile string) {
// 	outputFile = fmt.Sprintf("data/%s.mp4", outputFile)
// 	cmd := exec.Command("ffmpeg", "-y",
// 		"-i", vPath,
// 		"-c", "copy",
// 		"-map", "0:v:0",
// 		"-shortest",
// 		outputFile,
// 	)

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("\nОшибка при склейке: %v\n", err)
// 		return
// 	}

// 	os.Remove(vPath)

// 	fmt.Printf("\nГотово! Видео сохранено в: %s\n", outputFile)
// }

// func ffmpegRecodeMp4NoAudio(vPath, outputFile string) {
// 	outputFile = fmt.Sprintf("data/%s.mp4", outputFile)
// 	cmd := exec.Command("ffmpeg", "-y",
// 		"-i", vPath,
// 		"-c:v", "libx264", // Перекодируем видео
// 		"-crf", "18", // Качество 18 - почти без потерь
// 		"-preset", "fast",
// 		"-c:a", "copy",
// 		"-map", "0:v:0",
// 		"-shortest",
// 		outputFile,
// 	)

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("\nОшибка при склейке: %v\n", err)
// 		return
// 	}

// 	os.Remove(vPath)

// 	fmt.Printf("\nГотово! Видео сохранено в: %s\n", outputFile)
// }

// func ffmpegBuildMp3(aPath, outputFile string) {
// 	outputFile = fmt.Sprintf("data/%s.mp3", outputFile)
// 	cmd := exec.Command("ffmpeg", "-y",
// 		"-i", aPath,
// 		"-vn",
// 		"-c:a", "libmp3lame",
// 		"-q:a", "2", // Качество 2 - почти без потерь
// 		outputFile,
// 	)

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("\nОшибка при переводе в mp3: %v\n", err)
// 		return
// 	}

// 	os.Remove(aPath)

// 	fmt.Printf("\nГотово! аудио сохранено в: %s\n", outputFile)
// }
