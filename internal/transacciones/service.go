package internal

import "errors"

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	Search(id string) (Transaccion, error)
	Filter(mapEtiquetas, mapRelacionEtiquetas map[string]string) ([]Transaccion, error)
	Update(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	Delete(id int) (Transaccion, error)
	UpdateCodigoYMonto(id int, conTransaccion string, monto float64) (Transaccion, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (serv *service) GetAll() ([]Transaccion, error) {
	transacciones, err := serv.repository.getAll()
	if err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (serv *service) Store(codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	transaccion, err := serv.repository.Store(codTransaccion, moneda, monto, emisor, receptor, fechaTrans)
	if err != nil {
		return Transaccion{}, err
	}

	return transaccion, err
}

func (serv *service) Search(id string) (Transaccion, error) {
	return serv.repository.Search(id)
}

func (serv *service) Filter(mapEtiquetas, mapRelacionEtiquetas map[string]string) ([]Transaccion, error) {
	filtredTransac, err := serv.repository.Filter(mapEtiquetas, mapRelacionEtiquetas)
	if err != nil {
		return []Transaccion{}, err
	}
	if len(filtredTransac) == 0 {
		err = errors.New("no se encontrĂ³ ninguna transaccion para el filtro indicado")
	} else {
		err = nil
	}
	return filtredTransac, err
}

func (serv *service) Update(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	return serv.repository.Update(id, codTransaccion, moneda, monto, emisor, receptor, fechaTrans)
}

func (serv *service) Delete(id int) (Transaccion, error) {
	return serv.repository.Delete(id)
}

func (serv *service) UpdateCodigoYMonto(id int, conTransaccion string, monto float64) (Transaccion, error) {
	return serv.repository.UpdateCodigoYMonto(id, conTransaccion, monto)
}
