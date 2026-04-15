components {
  id: "8"
  component: "/components/bosses/8/8.script"
}
components {
  id: "spawn"
  component: "/components/particles/spawn.particlefx"
}
embedded_components {
  id: "bullet_factory"
  type: "factory"
  data: "prototype: \"/components/boss_bullets/red/red_bullet.go\"\n"
  "load_dynamically: true\n"
  "dynamic_prototype: true\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"boss\"\n"
  "mask: \"walls\"\n"
  "mask: \"player_bullets\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"esfera\"\n"
  "  }\n"
  "  data: 96.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"boss8\"\n"
  "material: \"/assets/shaders/sprite_color_add.material\"\n"
  "size {\n"
  "  x: 320.0\n"
  "  y: 180.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/core/core.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
  scale {
    x: 4.0
    y: 4.0
  }
}
