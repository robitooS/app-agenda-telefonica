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

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave({
      id: Number(id),
      nome,
      idade: Number(idade),
      telefones: telefones.map(t => ({ ...t, id_contato: Number(id) }))
    });
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
          </div>
          <div className="form-group">
            <label>Nome</label>
            <input 
              type="text" 
              value={nome} 
              onChange={(e) => setNome(e.target.value)} 
              required 
            />
          </div>
          <div className="form-group">
            <label>Idade</label>
            <input 
              type="number" 
              value={idade} 
              onChange={(e) => setIdade(e.target.value)} 
              required 
            />
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
                  value={tel.numero} 
                  onChange={(e) => updateTelefone(index, e.target.value)} 
                  required
                />
                <button type="button" onClick={() => removeTelefone(index)} className="btn-danger btn-icon">
                  <Trash2 size={16} />
                </button>
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
