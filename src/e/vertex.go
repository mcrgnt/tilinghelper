package e

// fragmentShaderSource = `
// 	#ifdef GL_ES
// 	precision mediump float;
// 	#endif

// 	uniform vec2 u_resolution;
// 	uniform vec2 u_mouse;
// 	uniform float u_time;

// 	void main() {
// 	    vec2 st = gl_FragCoord.xy/u_resolution.xy;
// 	    st.x *= u_resolution.x/u_resolution.y;

//     	vec3 color = vec3(0.);
//     	color = vec3(st.x,st.y,abs(sin(u_time)));

//     	gl_FragColor = vec4(color,1.0);
// 	}
// ` + "\x00"

// fragmentShaderSource = `
// #version 460
// out vec4 frag_colour;
// void main() {
// 	frag_colour = vec4(0.1, 0.5, 1, 1.0);
// }
// ` + "\x00"

// fragmentShaderSource = `
	// 	#version 460
	// 	in vec3 in_frag_colour;
	// 	out vec4 frag_colour;
	// 	void main() {
	// 		frag_colour = vec4(in_frag_colour, 1.0);
	// 	}
	// ` + "\x00"

	// fragmentShaderSource = `
	// void main() {
	// 	gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);
	// }
	// ` + "\x00"