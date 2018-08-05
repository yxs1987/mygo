package response_code

import "github.com/gin-gonic/gin"

func success(code int,msg string){
	return gin.H{"dsds"}
}

func err(code int,msg string)  {
	
}