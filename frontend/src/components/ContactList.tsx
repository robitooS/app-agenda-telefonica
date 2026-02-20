import React from 'react';
import { type Contato } from '../types';
import { Edit2, Trash2, Phone } from 'lucide-react';

interface ContactListProps {
  contacts: Contato[];
  onEdit: (contact: Contato) => void;
  onDelete: (id: number) => void;
}

export const ContactList: React.FC<ContactListProps> = ({ contacts, onEdit, onDelete }) => {
  return (
    <div className="contact-list-container">
      {contacts.length === 0 ? (
        <p className="no-results">Nenhum contato encontrado.</p>
      ) : (
        <table className="contact-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Nome</th>
              <th>Idade</th>
              <th>Telefones</th>
              <th>Ações</th>
            </tr>
          </thead>
          <tbody>
            {contacts.map((contact) => (
              <tr key={contact.id}>
                <td>{contact.id}</td>
                <td>{contact.nome}</td>
                <td>{contact.idade}</td>
                <td>
                  <div className="phone-list">
                    {contact.telefones?.map((t) => (
                      <div key={t.id} className="phone-tag">
                        <Phone size={12} /> {t.numero}
                      </div>
                    ))}
                  </div>
                </td>
                <td className="actions-cell">
                  <button onClick={() => onEdit(contact)} className="btn-icon btn-edit" title="Editar">
                    <Edit2 size={18} />
                  </button>
                  <button onClick={() => onDelete(contact.id)} className="btn-icon btn-delete" title="Excluir">
                    <Trash2 size={18} />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};
