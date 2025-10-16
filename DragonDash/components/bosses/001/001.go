components {
  id: "001"
  component: "/components/bosses/001/001.script"
}
embedded_components {
  id: "bullet_factory"
  type: "factory"
  data: "prototype: \"/components/boss_bullets/basic_green_bullet/basic_green_bullet.go\"\n"
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
  "mask: \"player_bullets\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1165.0\n"
  "      y: 359.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 117.5\n"
  "  data: 360.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"boss1\"\n"
  "material: \"/assets/shaders/sprite_color_add.material\"\n"
  "size {\n"
  "  x: 320.0\n"
  "  y: 180.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/levels/001/001.atlas\"\n"
  "}\n"
  ""
  position {
    x: 640.0
    y: 360.0
    z: 0.5
  }
  scale {
    x: 4.0
    y: 4.0
  }
}
