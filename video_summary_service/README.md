# Video Summary Service

This service exposes NemoStation/Marlin-2B video understanding over HTTP.

## Endpoints

- `GET /health`
- `POST /caption`: multipart video upload, returns Marlin dense caption output.
- `POST /find`: multipart video upload plus `event`, returns the grounded time span for the event.

## Quick Try

Start the service:

```bash
uvicorn video_summary_service.app:app --host 0.0.0.0 --port 8002
```

Caption a video:

```bash
curl -F file=@video.mp4 http://127.0.0.1:8002/caption
```

Find an event in a video:

```bash
curl -F file=@video.mp4 -F 'event=a person enters the room' http://127.0.0.1:8002/find
```

## Notes

- Configure the model under `video_summary` in `config.yaml`.
- The default model path is `./models/marlin`, matching the local model directory in this repository.
- Set `VIDEO_SUMMARY_MODEL=NemoStation/Marlin-2B` to load from Hugging Face instead.
