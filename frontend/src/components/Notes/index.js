import React, { useState, useEffect } from 'react';
import './styles.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEdit, faTrash, faSave, faCancel } from '@fortawesome/free-solid-svg-icons';
import api from '../../services/api';

function Notes() {
  const [notes, setNotes] = useState([]);
  const [newNote, setNewNote] = useState('');
  const [newTitle, setNewTitle] = useState('');
  const [editMode, setEditMode] = useState(null);
  const [editedContent, setEditedContent] = useState('');
  const [editedTitle, setEditedTitle] = useState('');

  useEffect(() => {
    api.get('/notes').then((response) => {
      setNotes(response.data.data.notes);
    })
    .catch(error => console.log(error));
  }, []);

  const addNote = () => {
    api
      .post('/notes', {
        title: newTitle,
        content: newNote,
      })
      .then((response) => {
        setNotes([...notes, response.data.data.note]);
        setNewTitle('');
        setNewNote('');
      })
      .catch(error => console.log(error));
  };

  const deleteNote = (id) => {
    api.delete(`/notes/${id}`).then(() => {
      const updatedNotes = notes.filter((note) => note.id !== id);
      setNotes(updatedNotes);
    })
    .catch(error => console.log(error));
  };

  const updateNote = (id) => {
    api
      .put(`/notes/${id}`, { 
        content: editedContent, 
        title: editedTitle 
      })
      .then(() => {
        const updatedNotes = [...notes];
        const noteIndex = updatedNotes.findIndex((note) => note.id === id);
        updatedNotes[noteIndex].title = editedTitle;
        updatedNotes[noteIndex].content = editedContent;
        setNotes(updatedNotes);
        setEditMode(null);
      })
      .catch(error => console.log(error));
  };

  return (
    <div className="notes-container">
      <h1 style={{color:"white"}}>Notes</h1>
      <div className="add-note">
        <input
          type="text"
          placeholder="Title..."
          value={newTitle}
          onChange={(e) => setNewTitle(e.target.value)}
        />
        <input
          type="text"
          placeholder="Note content..."
          value={newNote}
          onChange={(e) => setNewNote(e.target.value)}
        />
        <button onClick={addNote}>Add</button>
      </div>
      <div className="notes-list">
        {notes.map((note) => (
          <div key={note.id} className="note">
            <div className="note-content">
              {editMode === note.id ? (
                <div>
                  <input
                    type="text"
                    value={editedTitle}
                    onChange={(e) => setEditedTitle(e.target.value)}
                  />
                  <input
                    type="text"
                    value={editedContent}
                    onChange={(e) => setEditedContent(e.target.value)}
                  />
                  <button onClick={() => updateNote(note.id)} title={"Save"}>
                    <FontAwesomeIcon icon={faSave} />
                  </button>
                  <button onClick={() => setEditMode(null)} title={"Cancel edit"}>
                    <FontAwesomeIcon icon={faCancel} />
                  </button>
                </div>
              ) : (
                <div>
                  <h3>{note.title}</h3>
                  <p>{note.content}</p>
                </div>
              )}
            </div>
            <div className="note-date">{note.date}</div>
            <div className="note-actions">
              {editMode !== note.id ? (
                <div>
                  <button onClick={() => setEditMode(note.id)}  title={"Edit note"}>
                    <FontAwesomeIcon icon={faEdit} />
                  </button>
                </div>
              ) : null}
              <button className="delete-button" onClick={() => deleteNote(note.id)}  title={"Delete note"}>
                <FontAwesomeIcon icon={faTrash} />
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Notes;
