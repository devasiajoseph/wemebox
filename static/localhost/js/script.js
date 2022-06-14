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

  const video = document.querySelector('#video');
  function pic(){
    const canvas = document.querySelector('#canvas')
    canvas.width = video.videoWidth
    canvas.height = video.videoHeight
    canvas.getContext('2d').drawImage(video, 0, 0)
  }
  async function start(){
    const videoStream =  await navigator.mediaDevices.getUserMedia(constraints);
    video.srcObject = videoStream;
  }




  start();
  
