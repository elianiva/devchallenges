<script lang="ts">
  import Upload from "./components/Upload.svelte";
  import Loading from "./components/Loading.svelte";
  import Done from "./components/Done.svelte";
  import { isUploading, isFinished } from "./store";

  let imgSrc = "";

  async function upload(file: File) {
    isUploading.set(true);
    try {
      const formData = new FormData();
      formData.set("image", file);

      const response = await fetch("http://localhost:8000/upload", {
        method: "POST",
        body: formData,
      });
      const data = await response.json();

      imgSrc = data.file_url;
      isUploading.set(false);
      isFinished.set(true);
    } catch (err) {
      console.error(err);
    }
  }

  function dropHandler(event: DragEvent) {
    if (event.dataTransfer.items) {
      upload(event.dataTransfer.items[0].getAsFile());
    } else {
      upload(event.dataTransfer.files[0]);
    }
  }

  let fileInput: HTMLInputElement;
</script>

<div class="container">
  <div class="box">
    {#if $isUploading && !$isFinished}
      <Loading />
    {:else if !$isUploading && $isFinished}
      <Done src={imgSrc} />
    {:else}
      <Upload {dropHandler} {fileInput} {upload} />
    {/if}
  </div>
</div>

<style>
  .container {
    width: 100%;
    min-height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 1rem;
  }

  .box {
    background-color: #ffffff;
    border-radius: 0.5rem;
    box-shadow: 0 0.25rem 0.75rem rgba(0, 0, 0, 0.1);
    padding: 2rem;
    text-align: center;
  }
</style>
