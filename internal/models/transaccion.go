package models

type Ciudad struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	NombrePais string `json:"nombre_pais"`
}

type Transaccion struct {
	CodTransaccion string  `json:"cod_transaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         string  `json:"emisor"`
	Receptor       string  `json:"receptor"`
	FechaTrans     string  `json:"fecha_trans"`
	Ciudad         Ciudad  `json:"ciudad"`
}
