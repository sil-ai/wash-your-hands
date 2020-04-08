#!/bin/bash

for f in poster_text/*
do
	pango-view $f --no-display --dpi=400 --font=unifont.ttf --output poster_text_render/$(basename $f).png
done
