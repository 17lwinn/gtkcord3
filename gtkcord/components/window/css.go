package window

import (
	"github.com/diamondburned/gtkcord3/internal/log"
	"github.com/gotk3/gotk3/gtk"
)

const CSS = `
	@define-color color_discord rgb(114, 137, 218);
	@define-color color_pinged  rgb(240, 71, 71);

	button.flat {
		padding: 0;
		min-height: 0;
	}

	.codeblock {
		background-color: alpha(@theme_bg_color, 0.25);
		padding: 8px;
	}

	.status {
		background-color: white;
		border: 2px solid #FFFFFF;
		border-radius: 9999px;
	}
	.guilds     .status { border-width: 3px; }
	.popup-user .status { border-width: 4px; }

	.status.online {
		border-color: #43B581;
	}
	.status.busy {
		border-color: #F04747;
	}
	.status.idle {
		border-color: #FAA61A;
	}
	.status.offline {
		border-color: #747F8D;
	}
	.status.unknown {
		border-color: #FFFFFF;
	}

	.quickswitcher list { background-color: transparent; }

	headerbar { padding: 0; }
	headerbar button { box-shadow: none; }
	textview, textview > text { background-color: transparent; }

	.guilds, .channels, .dmchannels, .members {
		background-color: @theme_bg_color;
	}

	.messagecontainer {
		background-color: @theme_base_color;
	}
	.messages {
		background-color: @theme_base_color;
	}
	.messages > row {
		padding: 0;
	}

	.message {
		border-left: 2px solid transparent;
	}
	.message.mentioned {
		border-left: 2px solid rgb(250, 166, 26);
		background-color: rgba(250, 166, 26, 0.05);
	}
	.messages > row       .message.condensed .timestamp {
		opacity: 0;
	}
	.messages > row:hover .message.condensed .timestamp {
		opacity: 1;
	}

	.guilds > * {
		padding-left: 0;
		padding-right: 0;
	}
	.guild-folder.unread {
		background-color: alpha(@theme_selected_bg_color, 0.15);
	}
	.guild-folder.pinged {
		background-color: alpha(@color_pinged, 0.15);
	}
	.guild-folder:selected {
		border-top: 5px solid alpha(@theme_selected_bg_color, 0.5);
	}
	.guild-folder:selected list {
		border-bottom: 5px solid alpha(@theme_selected_bg_color, 0.5);
	}
	.guild:selected {
		background-color: alpha(@theme_selected_bg_color, 0.5);
	}
	.guild:hover image {
		-gtk-icon-effect: highlight;
	}

	.emojiview list, .emojiview searchbar box {
		background-color: @theme_base_color;
	}
	.emojiview searchbar entry {
		background-color: transparent;
	}
	.emojiview button {
		padding: 4px;
		border-radius: 0;
	}
	.emojiview button:hover {
		background-image: none;
	}
	.emojiview flowbox {
		background-color: shade(@theme_bg_color, 0.95);
		padding: 4px 4px;
	}
	.emojiview row {
		padding: 0;
	}

	.user-info, .popup-grid > *:not(.popup-user):not(.activity) {
		background-color: shade(@theme_base_color, 0.9);
		color: @theme_fg_color;
	}
	.popup-grid > .activity {
		background-color: rgba(0, 0, 0, 0.12);
	}

	.user-info, popover {
		padding: 0;
		margin:  0;
	}
	.user-info separator {
		margin-top: 0;
	}
	.user-info.spotify {
		background-color: #1db954;
		color: white;
	}
	.user-info.twitch {
		background-color: rgb(89, 54, 149);
		color: white;
	}
	.user-info.game {
		background-color: #7289da;
		color: white;
	}

	.message-input {
		background-image: linear-gradient(transparent, rgba(10, 10, 10, 0.3) 60px);
		transition-property: background-image;
		transition: 75ms background-image linear;
	}
	.message-input.editing {
		background-image: linear-gradient(transparent, alpha(@color_discord, 0.3) 60px);
	}

	.message-input .completer {
		background-color: transparent;
	}

	.message-input button {
		background: none;
		box-shadow: none;
		border: none;
		opacity: 0.65;
	}
	.message-input button:hover {
		opacity: 1;
	}
	
	.guild > image, .dmbutton > image {
	    border-radius: 50%;
	}
	.guild.unread > image {
		border: 2px solid @theme_fg_color;
		padding: 2px;
	}
	.guild.pinged > image, .dmbutton.pinged > image {
		border: 2px solid rgb(240, 71, 71);
		padding: 2px;
	}

	.channel label, .category.muted label {
		opacity: 0.5;
	}
	.channel.muted label {
		opacity: 0.25;
	}
	.channel.unread label, .opacity.pinged label {
		opacity: 1;
	}
	.channel.pinged {
		color: @color_pinged;
		background-color: alpha(@color_pinged, 0.15);
	}
	.dmchannel.pinged {
		background-color: alpha(@color_pinged, 0.15);
	}

	.member-section list {
		background-color: transparent;
	}

	.reactions > flowboxchild {
		margin: 0;
		padding: 1px;
	}
	button.reaction:checked {
		background-color: alpha(@color_discord, 0.2);
		color: @color_discord;
	}
`

var (
	CustomCSS string // raw CSS, once
	FileCSS   string // path
)

// I don't like this:
// list row:selected { box-shadow: inset 2px 0 0 0 white; }

func initCSS() {
	s := Window.Screen

	stock, _ := gtk.CssProviderNew()
	if err := stock.LoadFromData(CSS); err != nil {
		log.Fatalln("Failed to parse stock CSS:", err)
	}

	gtk.AddProviderForScreen(
		s, stock,
		uint(gtk.STYLE_PROVIDER_PRIORITY_APPLICATION),
	)

	// Add env var CSS:
	env, _ := gtk.CssProviderNew()
	if err := env.LoadFromData(CustomCSS); err != nil {
		log.Errorln("Failed to parse env var custom CSS:", err)
	}

	gtk.AddProviderForScreen(
		s, env,
		uint(gtk.STYLE_PROVIDER_PRIORITY_USER),
	)
}

func ReloadCSS() {
	s := Window.Screen

	// Replace file CSS:
	if Window.fileCSS != nil {
		gtk.RemoveProviderForScreen(s, Window.fileCSS)
	}

	file, _ := gtk.CssProviderNew()
	if err := file.LoadFromPath(FileCSS); err != nil {
		log.Errorln("Failed to parse file in "+FileCSS+":", err)
		return
	}

	Window.fileCSS = file

	gtk.AddProviderForScreen(
		s, file,
		uint(gtk.STYLE_PROVIDER_PRIORITY_USER),
	)
}
