#version 410

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

in vec3 vert;
in vec3 norm;

out vec3 color;

void main() {
	// Set vertex position
	gl_Position = projection * view * model * vec4(vert, 1);

	// Set vertex color
	vec3 light = vec3(-1.0, 1.5, 0.85);
	vec3 base_color = vec3(0.2, 0.5, 0.7);
	float d = (dot(light, norm) + 1.0) * 0.1;
	color = base_color + vec3(d, d, d);
}
