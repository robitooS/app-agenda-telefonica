import axios from 'axios';
import { type Contato } from '../types';

const api = axios.create({
  baseURL: 'http://localhost:8080',
});

export const contactService = {
  list: (nome?: string, numero?: string) => 
    api.get<Contato[]>('/contatos', { params: { nome, numero } }),
  
  get: (id: number) => 
    api.get<Contato>(`/contatos/${id}`),
  
  create: (contato: Contato) => 
    api.post<Contato>('/contatos', contato),
  
  update: (id: number, contato: Contato) => 
    api.put<Contato>(`/contatos/${id}`, contato),
  
  delete: (id: number) => 
    api.delete(`/contatos/${id}`),
};
