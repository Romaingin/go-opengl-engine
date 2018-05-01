#version 410

in vec3 color;
in vec2 UV;

out vec4 frag_colour;

uniform sampler2D tex;

void main() {
	// frag_colour = vec4(color, 1.0);
	frag_colour = texture( tex, UV ).rgba;
}
