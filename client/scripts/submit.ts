import axios from 'axios';
import { errorToastify, successToastify } from './lib/toast';

const form = document.getElementById('form');

form.onsubmit = async e => {
  e.preventDefault();

  const webhookUrl = (<HTMLFormElement>document.getElementById('webhook-input')).value as string;
  if (!(webhookUrl.includes('discordapp.com') || webhookUrl.includes('hooks.slack.com')))
    return errorToastify('웹훅 URL이 정확한지 확인해주세요');

  (<HTMLButtonElement>document.getElementById('webhook-submit')).disabled = true;
  
  try {
    await axios.post('/api/subscribe', {
      url: webhookUrl,
    });

    successToastify('구독이 완료되었습니다!', 10000);
  } catch (e) {
    errorToastify('웹훅 URL이 정확한지 확인해주세요');
  }

  (<HTMLButtonElement>document.getElementById('webhook-submit')).disabled = false;

}

