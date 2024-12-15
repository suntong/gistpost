package gistpost

type CreateCommand struct {
	Public bool `short:"p" long:"pub" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
}

type FolderCommand struct {
	Dir    string `short:"D" long:"dir" description:"Directory to upload as gist*" required:"true"`
	Public bool   `short:"p" long:"pub" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
	Extra  bool   `short:"e" long:"extra" env:"GISTPOST_EXTRA" description:"Extra files will be added to gist for better name/doc."`
}

type UpdateCommand struct {
	GistID string `short:"g" long:"id" env:"GISTPOST_GISTID" description:"Existing GH gist id*" required:"true"`
}

var (
	Opts OptsT
)
