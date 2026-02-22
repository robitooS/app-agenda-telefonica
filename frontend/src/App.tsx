import { useState, useEffect } from 'react';
import { ContactList } from './components/ContactList';
import { ContactForm } from './components/ContactForm';
import { contactService } from './services/api';
import { type Contato } from './types';
import { Search, UserPlus } from 'lucide-react';
import './App.css';

function App() {
  const [contacts, setContacts] = useState<Contato[]>([
    {
      id: 1,
      nome: "Contato de Exemplo",
      idade: 30,
      telefones: [{ id: 1, numero: "(11) 99999-9999", id_contato: 1 }]
    }
  ]);
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [editingContact, setEditingContact] = useState<Contato | null>(null);
  const [searchName, setSearchName] = useState('');
  const [searchPhone, setSearchPhone] = useState('');
  const [loading, setLoading] = useState(false);

  const fetchContacts = async () => {
    setLoading(true);
    try {
      const response = await contactService.list(searchName, searchPhone);
      setContacts(response.data || []); // Garante que seja um array, mesmo que response.data seja null/undefined
    } catch (error) {
      console.warn('Backend offline ou erro ao buscar contatos, usando array vazio.');
      setContacts([]); // Define como array vazio em caso de erro
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchContacts();
  }, []);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    fetchContacts();
  };

  const handleSaveContact = async (contact: Contato) => {
    try {
      if (editingContact) {
        await contactService.update(contact.id, contact);
      } else {
        await contactService.create(contact);
      }
      setIsFormOpen(false);
      setEditingContact(null);
      fetchContacts();
    } catch (error: any) {
      console.error('Erro ao salvar contato:', error);
      alert(error.response?.data?.error || 'Erro ao salvar contato.');
    }
  };

  const handleDeleteContact = async (id: number) => {
    if (window.confirm('Tem certeza que deseja excluir este contato?')) {
      try {
        await contactService.delete(id);
        fetchContacts();
      } catch (error) {
        console.error('Erro ao excluir contato:', error);
        alert('Erro ao excluir contato.');
      }
    }
  };

  const openCreateForm = () => {
    setEditingContact(null);
    setIsFormOpen(true);
  };

  const openEditForm = (contact: Contato) => {
    setEditingContact(contact);
    setIsFormOpen(true);
  };

  return (
    <div className="app-container">
      <header className="app-header">
        <h1>Agenda Telefônica</h1>
        <button onClick={openCreateForm} className="btn-primary">
          <UserPlus size={20} /> Novo Contato
        </button>
      </header>

      <main className="app-content">
        <section className="search-section">
          <form onSubmit={handleSearch} className="search-form">
            <div className="search-group">
              <label>Nome</label>
              <input 
                type="text" 
                placeholder="Pesquisar por nome..." 
                value={searchName}
                onChange={(e) => setSearchName(e.target.value)}
              />
            </div>
            <div className="search-group">
              <label>Telefone</label>
              <input 
                type="text" 
                placeholder="Pesquisar por número..." 
                value={searchPhone}
                onChange={(e) => setSearchPhone(e.target.value)}
              />
            </div>
            <button type="submit" className="btn-secondary search-btn">
              <Search size={20} /> Pesquisar
            </button>
          </form>
        </section>

        {loading ? (
          <div className="loading">Carregando...</div>
        ) : (
          <ContactList 
            contacts={contacts} 
            onEdit={openEditForm} 
            onDelete={handleDeleteContact} 
          />
        )}
      </main>

      {isFormOpen && (
        <ContactForm 
          contact={editingContact} 
          onSave={handleSaveContact} 
          onCancel={() => setIsFormOpen(false)} 
        />
      )}
    </div>
  );
}

export default App;
