package media

import (
	"os"
	"os/exec"
	"io"
	"log"
)

func PlayMedia() {

	// Abre el archivo de video (o cualquier otra fuente)
	videoFile, err := os.Open("sample.mp4")
	if err != nil {
		log.Fatal("Error al abrir el archivo de video:", err)
	}
	defer videoFile.Close()

	// Configurar FFplay para leer desde stdin
	cmd := exec.Command("ffplay",
		"-i", "pipe:0", // Leer de stdin
		"-x", "1024", // Ancho de ventana
		"-y", "720", // Alto de ventana
		"-autoexit",                               // Salir cuando termine el video
		"-window_title", "LibreRemotePlay Stream") // Título de la ventana

	// Conecta stdin de FFplay
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal("Error al crear pipe de entrada:", err)
	}

	// Opcional: conectar la salida estándar para ver mensajes de FFplay
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Iniciar el proceso FFplay
	if err := cmd.Start(); err != nil {
		log.Fatal("Error al iniciar FFplay:", err)
	}

	// Enviar el contenido del archivo a FFplay
	go func() {
		defer stdin.Close()
		if _, err := io.Copy(stdin, videoFile); err != nil {
			log.Fatal("Error al enviar datos a FFplay:", err)
		}
	}()

	// Esperar a que FFplay termine
	if err := cmd.Wait(); err != nil {
		log.Fatal("FFplay terminó con error:", err)
	}
}