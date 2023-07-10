# Jellyfin Prometheus Exporter
Collect Prometheus Metrics about a Jellyfin Instance


## Run with Docker

```
docker run -e TOKEN=c817af2a9b484317afcdb2f82567709b \
  -e JF_URL=https://demo.jellyfin.org/stable/ \
  -p 8080:8080 \
  ghcr.io/stefanabl/jellyfin-prometheus-exporter
```

## Configuration

Environment Variables

| Name         |          Example           |                                                      Explanation |
|--------------|:--------------------------:|-----------------------------------------------------------------:|
| JF_URL | https://demo.jellyfin.org/ | Base URL of the Jellyfin Instance<br/>**Note the traling slash** |
| TOKEN      |      REALLYLONGSTRING      |                                             Authentication Token |

Command Line Flags

`--listen-address=:8080` (Default is ":8080")


## Example Output

```
# HELP jellyfin_active_streams_direct_count The number of streams which are currently being direct streams
# TYPE jellyfin_active_streams_direct_count gauge
jellyfin_active_streams_direct_count 0
# HELP jellyfin_active_streams_transcode_count The number of streams which are currently being transcoded
# TYPE jellyfin_active_streams_transcode_count gauge
jellyfin_active_streams_transcode_count 0
# HELP jellyfin_items_count Count of Media Items label denotes type
# TYPE jellyfin_items_count gauge
jellyfin_items_count{type="AlbumCount"} 0
jellyfin_items_count{type="ArtistCount"} 0
jellyfin_items_count{type="BookCount"} 0
jellyfin_items_count{type="BoxSetCount"} 3
jellyfin_items_count{type="EpisodeCount"} 44
jellyfin_items_count{type="ItemCount"} 0
jellyfin_items_count{type="MovieCount"} 10
jellyfin_items_count{type="MusicVideoCount"} 0
jellyfin_items_count{type="ProgramCount"} 0
jellyfin_items_count{type="SeriesCount"} 2
jellyfin_items_count{type="SongCount"} 0
jellyfin_items_count{type="TrailerCount"} 0
# HELP jellyfin_streams_bandwidth_bits The total bandwidth currently being streamed
# TYPE jellyfin_streams_bandwidth_bits gauge
jellyfin_streams_bandwidth_bits 0
```

## Additional Notes

The exporter also exposes the endpoint `/rel-sizes`, which computes the "compression ratio" for **every item** on your server.
This operation can take very long to complete and cause high load on Jellyfin.
The relative size is computed as: File Size / ( Duration * Resolution ).

It can be used in combination with the [JSON API Grafana Plugin](https://grafana.com/grafana/plugins/marcusolsson-json-datasource/) to visualize which of your files are compressed inefficiently.

An example response in JSON:
```josn
[
  {
    "name": "Night of the Living Dead",
    "path": "/media/Movies/Archive.org/Night of the Living Dead (1968)/Night of the Living Dead (1968).mp4",
    "resolution": "640x480",
    "size_bytes": 596645312,
    "length_seconds": 5717,
    "rel_size": 0.3397244694186928,
    "codec": "h264"
  }
]
```