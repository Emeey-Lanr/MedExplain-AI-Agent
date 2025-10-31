package main
import( "github.com/gin-gonic/gin"
"github.com/joho/godotenv"
"log"
)



func main ()  {

	r := gin.Default()

	if err := godotenv.Load(); err != nil{
		log.Fatal("Error loading.env:", err) // change to println in production for your app not crash
	}

	


	r.Run(":8080")
}