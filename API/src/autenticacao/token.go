package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões de usuario
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey)) //chave secret que assina o token
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token inválido")
}

// ExtrairUsuarioID retorna o usuarioId que está salvo no Token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	fmt.Printf("Token recebido: %s\n", tokenString) // Depuração do token

	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		fmt.Println("Erro ao fazer parse do token:", erro)
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Permissões extraídas do token: %+v\n", permissoes) // Depuração

		valor, existe := permissoes["usuarioId"]
		if !existe {
			fmt.Println("Campo 'usuarioID' não encontrado nas permissões")
			return 0, errors.New("Token inválido: 'usuarioID' não encontrado")
		}

		// Verificando se o tipo do valor é compatível com float64
		usuarioIDFloat, ok := valor.(float64)
		if !ok {
			fmt.Printf("Tipo inesperado para 'usuarioID': %T\n", valor)
			return 0, errors.New("Tipo de 'usuarioID' inválido no token")
		}

		// Convertendo para uint64
		usuarioID := uint64(usuarioIDFloat)
		fmt.Printf("usuarioID extraído do token: %d\n", usuarioID)
		return usuarioID, nil
	}

	fmt.Println("Token inválido ou permissões ausentes")
	return 0, errors.New("Token inválido")
}

// extrairToken extrai o token da requisição, pois ele vem com um Bearer - antes
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
