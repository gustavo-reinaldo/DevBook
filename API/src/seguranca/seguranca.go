package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

// VerificarSenha compara uma senha com um hash e retorna se são iguais
func VerificarSenha(senhaComHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))
}
