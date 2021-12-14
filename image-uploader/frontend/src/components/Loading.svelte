<h1 class="loading__title">Uploading...</h1>
<div class="loading__progress">
  <div class="loading__progress-thumb" />
</div>

<style>
  .loading__title {
    text-align: left;
    font-family: "Poppins", sans-serif;
    font-size: 1.25rem;
    font-weight: 500;
    color: #4f4f4f;
    margin-bottom: 1rem;
  }

  .loading__progress {
    width: 20rem;
    height: 0.5rem;
    background-color: #f2f2f2;
    border-radius: 0.25rem;
    overflow: hidden;
  }

  .loading__progress-thumb {
    width: 6rem;
    height: 0.5rem;
    border-radius: 0.25rem;
    background-color: #2f80ed;
    animation: move 1.75s ease-out infinite;
  }

  @keyframes move {
    0% {
      transform: translateX(-6rem);
    }
    100% {
      transform: translateX(26rem);
    }
  }
</style>
