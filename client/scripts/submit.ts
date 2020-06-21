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
    if (e.response.status === 409)
      errorToastify('이미 등록 된 웹훅 URL 이에요');
    else if (e.response.status === 403)
      errorToastify('웹훅 URL이 정확한지 확인해주세요');
    else
      errorToastify('알 수 없는 오류가 발생했어요.');
  }

  (<HTMLButtonElement>document.getElementById('webhook-submit')).disabled = false;

}

