export interface Telefone {
  id_contato?: number;
  id: number;
  numero: string;
}

export interface Contato {
  id: number;
  nome: string;
  idade: number;
  telefones: Telefone[];
}
