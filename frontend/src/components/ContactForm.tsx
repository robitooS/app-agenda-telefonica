import React, { useState, useEffect } from 'react';
import { type Contato, type Telefone } from '../types';
import { Plus, Trash2, X } from 'lucide-react';

interface ContactFormProps {
  contact?: Contato | null;
  onSave: (contact: Contato) => void;
  onCancel: () => void;
}

export const ContactForm: React.FC<ContactFormProps> = ({ contact, onSave, onCancel }) => {
  const [nome, setNome] = useState('');
  const [idade, setIdade] = useState('');
  const [telefones, setTelefones] = useState<Telefone[]>([]);
  const [id, setId] = useState('');
  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  useEffect(() => {
    if (contact) {
      setId(contact.id.toString());
      setNome(contact.nome);
      setIdade(contact.idade.toString());
      setTelefones(contact.telefones || []);
    } else {
      setId('');
      setNome('');
      setIdade('');
      setTelefones([]);
    }
    setErrors({}); // Limpar erros ao abrir/editar novo contato
  }, [contact]);

  const addTelefone = () => {
    setTelefones([...telefones, { id: Date.now(), numero: '' }]);
  };

  const removeTelefone = (index: number) => {
    setTelefones(telefones.filter((_, i) => i !== index));
  };

  const updateTelefone = (index: number, numero: string) => {
    const newTelefones = [...telefones];
    newTelefones[index].numero = numero;
    setTelefones(newTelefones);
  };

  const validate = () => {
    const newErrors: { [key: string]: string } = {};

    if (!nome || nome.trim().length < 2) {
      newErrors.nome = 'Nome deve ter no mínimo 2 caracteres.';
    }
    if (!idade || Number(idade) < 0) {
      newErrors.idade = 'Idade não pode ser negativa.';
    }
    if (!id || Number(id) <= 0 || !Number.isInteger(Number(id))) {
        newErrors.id = 'ID deve ser um número inteiro positivo.';
    }

    // Validação para telefones
    telefones.forEach((tel, index) => {
      if (!tel.numero || tel.numero.replace(/\D/g, '').length < 10) { // Mínimo de 10 dígitos para telefone
        newErrors[`telefone-${index}`] = 'Telefone inválido (mínimo 10 dígitos).';
      }
    });

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!validate()) {
      return;
    }
    onSave({
      id: Number(id),
      nome,
      idade: Number(idade),
      telefones: telefones.map(t => ({ ...t, id_contato: Number(id) }))
    });
  };

  const formatPhoneNumber = (value: string) => {
    if (!value) return value;
    const phoneNumber = value.replace(/\D/g, ''); // Remove tudo que não é dígito
    const length = phoneNumber.length;

    if (length <= 2) return `(${phoneNumber}`;
    if (length <= 6) return `(${phoneNumber.slice(0, 2)}) ${phoneNumber.slice(2)}`;
    if (length <= 10) return `(${phoneNumber.slice(0, 2)}) ${phoneNumber.slice(2, 6)}-${phoneNumber.slice(6)}`;
    return `(${phoneNumber.slice(0, 2)}) ${phoneNumber.slice(2, 7)}-${phoneNumber.slice(7, 11)}`;
  };
  
  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <div className="modal-header">
          <h2>{contact ? 'Alterar Contato' : 'Novo Contato'}</h2>
          <button onClick={onCancel} className="btn-icon"><X /></button>
        </div>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>ID (Numérico)</label>
            <input 
              type="number" 
              value={id} 
              onChange={(e) => setId(e.target.value)} 
              required 
              disabled={!!contact}
            />
            {errors.id && <span className="error-message">{errors.id}</span>}
          </div>
          <div className="form-group">
            <label>Nome</label>
            <input 
              type="text" 
              value={nome} 
              onChange={(e) => setNome(e.target.value)} 
              required 
            />
            {errors.nome && <span className="error-message">{errors.nome}</span>}
          </div>
          <div className="form-group">
            <label>Idade</label>
            <input 
              type="number" 
              value={idade} 
              onChange={(e) => setIdade(e.target.value)} 
              required 
            />
            {errors.idade && <span className="error-message">{errors.idade}</span>}
          </div>
          
          <div className="phones-section">
            <div className="section-header">
              <h3>Telefones</h3>
              <button type="button" onClick={addTelefone} className="btn-secondary btn-sm">
                <Plus size={16} /> Adicionar
              </button>
            </div>
            {telefones.map((tel, index) => (
              <div key={tel.id} className="phone-input-group">
                <input 
                  type="text" 
                  placeholder="Número" 
                  value={formatPhoneNumber(tel.numero)} 
                  onChange={(e) => updateTelefone(index, e.target.value)} 
                  required
                />
                <button type="button" onClick={() => removeTelefone(index)} className="btn-danger btn-icon">
                  <Trash2 size={16} />
                </button>
                {errors[`telefone-${index}`] && <span className="error-message">{errors[`telefone-${index}`]}</span>}
              </div>
            ))}
          </div>

          <div className="form-actions">
            <button type="button" onClick={onCancel} className="btn-secondary">Cancelar</button>
            <button type="submit" className="btn-primary">Salvar</button>
          </div>
        </form>
      </div>
    </div>
  );
};
