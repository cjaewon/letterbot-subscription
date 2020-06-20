import Toastify from 'toastify-js';
import 'toastify-js/src/toastify.css';

export function successToastify(text: string, duration: Number = 3000) {
  const toast = Toastify({
    text,
    duration,
    backgroundColor: '#38A169',
    position: 'center',
    close: true,
  }) as any;
  
  toast.showToast();
}

export function errorToastify(text: string, duration: Number = 3000) {
  const toast = Toastify({
    text,
    duration,
    backgroundColor: '#F56565',
    position: 'center',
    close: true,
  }) as any;
  
  toast.showToast();
}