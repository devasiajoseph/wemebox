if ('mediaDevices' in navigator && 'getUserMedia' in navigator.mediaDevices) {
    console.log("Let's get this party started")
  };


  const constraints = {
    video: {
      width: {
        min: 1280,
        ideal: 1920,
        max: 2560,
      },
      height: {
        min: 720,
        ideal: 1080,
        max: 1440,
      },
    },
  }

  const video = document.querySelector('#video')
  async function start(){
    const videoStream =  await navigator.mediaDevices.getUserMedia(constraints);
    video.srcObject = videoStream;
  }


  start();
  
