const audioPlayer = document.querySelector("[data-audioPlayer]");
const pronounciationAudio = document.querySelector("[data-pronounciation]");

audioPlayer.addEventListener("click", () => {
  console.log("CLicked");
  pronounciationAudio.play();
});
