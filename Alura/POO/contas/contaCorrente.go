package contas

import "awesomeProject/banco/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar{
		c.saldo -= valorDoSaque
		return "Saque autorizado"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar (valorDoDeposito float64) (string,float64) {
	if valorDoDeposito > 0{
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!",c.saldo
	} else{
		return "Esse valor não pode ser depositado!", c.saldo
	}
}
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0{
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64  {
	return c.saldo
}
