const form = document.getElementById('form');

form.onsubmit = async e => {
  e.preventDefault();

  const webhookUrl = (<HTMLFormElement>document.getElementById('webhook-input')).value as string;
  if (!webhookUrl.includes('discordapp.com') || !webhookUrl.includes('hooks.slack.com')) return; //TODO: print error;

  
}