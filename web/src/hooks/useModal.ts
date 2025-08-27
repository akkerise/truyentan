import { useContext } from 'react';
import { UIContext } from '../context/UIContext';

export default function useModal(name: string) {
  const context = useContext(UIContext);
  if (!context) {
    throw new Error('useModal must be used within a UIProvider');
  }
  const { modals, openModal, closeModal } = context;
  const isOpen = !!modals[name];
  return {
    isOpen,
    open: () => openModal(name),
    close: () => closeModal(name),
  };
}
