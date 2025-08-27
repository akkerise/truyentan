import { Suspense, lazy } from 'react';
import useModal from '../hooks/useModal';
import Modal from '../components/common/Modal';

const Settings = lazy(() => import('./Settings'));

export default function Home() {
  const { isOpen, open, close } = useModal();

  return (
    <div>
      Home
      <button onClick={open}>Open Settings</button>
      <Suspense fallback={null}>
        <Modal isOpen={isOpen} onClose={close}>
          <Settings />
        </Modal>
      </Suspense>
    </div>
  );
}
