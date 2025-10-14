pub enum Token {
    Constant(f64),
    Operator(Operator),
    Bracket(Side)
}

pub enum Side {Left, Right}

macro_rules! define_operators {
    ($vis:vis $enum_name:ident {$( $name:ident $(,)? )*}) => {
        $vis enum $enum_name {
            $($name,)*
        }
    };
}

define_operators!(pub Operator {
    Addition,
    Negation,
    Multiplication,
    Division,
    Exponentiation,
    RootExtraction
});
