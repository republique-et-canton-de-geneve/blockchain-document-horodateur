%% This file implements a poster template for the 2014 Radboud 
%% University corporate style. 
%% 
%% For comments, questions, and suggestions contact me at
%% l.onrust@let.ru.nl 
%%
%% You can distribute and edit the files as you wish, but I'd
%% love to hear of any changes. Also, if you let me know that
%% you are using the template, I can keep you up-to-date on
%% future changes.

\ProvidesPackage{beamerthemeruhuisstijlposter}

\LoadClass[final,hyperref={pdfpagelabels=false}]{beamer}


\newlength{\columnheight}
\setlength{\columnheight}{30cm}

\RequirePackage{ifthen}
\RequirePackage[size=customflyer,scale=1.0,debug]{beamerposter}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\definecolor{ruhuisstijlrood}{RGB}{230,51,18}
\definecolor{grey}{RGB}{237,237,237}
\definecolor{darkgrey}{RGB}{43,43,40}
\definecolor{lightgrey}{RGB}{210,210,210}
% \definecolor{ruhuisstijlrood}{RGB}{0,200,0}
\definecolor{ruhuisstijlzwart}{rgb}{0,0,0}
\definecolor{ruhuisstijlwit}{rgb}{0.98,0.98,0.98}

\RequirePackage{tcolorbox}
\tcbuselibrary{skins,hooks}
\usetikzlibrary{shadows}

\tcbset{colframe=structure,fonttitle=\bfseries,beamer}

\newcommand{\roundedcorners}{false}
\DeclareOptionBeamer{roundedcorners}{\renewcommand{\roundedcorners}{#1}}

\newcommand{\titleposition}{right}
\DeclareOptionBeamer{titleposition}{\renewcommand{\titleposition}{#1}}

\ProcessOptionsBeamer

\mode<presentation>

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\setbeamercolor{headline}{fg=ruhuisstijlzwart,bg=lightgrey}
\setbeamercolor{footline}{fg=ruhuisstijlwit, bg=darkgrey}
\setbeamerfont{footline}{size=\large,series=\bf}
\setbeamercolor{title in headline}{fg=ruhuisstijlzwart}
\setbeamercolor{author in headline}{fg=ruhuisstijlzwart}
\setbeamercolor{institute in headline}{fg=ruhuisstijlzwart}

\setbeamercolor{author in head/foot}{fg=ruhuisstijlrood, bg=black}
\setbeamercolor{title in head/foot}{fg=ruhuisstijlrood, bg=black}

\setbeamerfont{normal text}{size=\fontsize{5cm}{1em},series=\sl}
\setbeamercolor*{normal text}{fg=ruhuisstijlzwart, bg=ruhuisstijlwit}
\setbeamercolor*{block body}{bg=ruhuisstijlwit,fg=ruhuisstijlzwart}
\setbeamercolor*{block title}{fg=ruhuisstijlzwart,bg=grey}
\setbeamerfont*{block title}{size=\large,series=\bf}


\setbeamercolor{alerted text}{fg=ruhuisstijlzwart}

\setbeamercolor{example text}{fg=ruhuisstijlzwart}
\setbeamercolor{structure}{fg=ruhuisstijlzwart}

\setbeamertemplate{itemize items}[triangle]
\setbeamertemplate{navigation symbols}{} 

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
\setbeamertemplate{block begin}{
  \vskip.75ex
  \begin{beamercolorbox}[rounded=\roundedcorners,dp=0.5ex,center,leftskip=-1em,colsep*=.75ex]{block title}%
    \usebeamerfont*{block title}%
    {\strut\phantom{GgM}\Large\insertblocktitle\strut\par}% phantom because of baseline problem
  \end{beamercolorbox}%
  {\ifbeamercolorempty[bg]{block body}{}{\nointerlineskip\vskip-0.5pt}}%
  \usebeamerfont{block body}\large%
  \begin{beamercolorbox}[leftskip=1em,colsep*=.75ex,sep=0.5ex,vmode]{block body}%
    \ifbeamercolorempty[bg]{block body}{\vskip-.25ex}{\vskip-.75ex}\vbox{}%
  }
  \setbeamertemplate{block end}{
  \end{beamercolorbox}
}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
\setbeamertemplate{headline}{  
  \leavevmode

  \begin{beamercolorbox}[wd=\paperwidth]{headline}
    \begin{columns}[T]
      \begin{column}{.02\paperwidth}
      \end{column}
      \begin{column}{.95\paperwidth}
        \vskip4ex
        \raggedleft
        {\ifthenelse{\equal{\titleposition}{right}}{%
        	\raggedleft\usebeamercolor{title in headline}{\color{fg}{\Huge\bfseries\inserttitle\par}}
        }{
          \ifthenelse{\equal{\titleposition}{left}}{%
              \raggedright\usebeamercolor{title in headline}{\color{fg}{\Huge\bfseries\inserttitle\par}}
          }
          {\centering
          \usebeamercolor{title in headline}{\color{fg}{\Huge\bfseries\inserttitle\par}}
          }
        }}
        \usebeamercolor{author in headline}{\color{fg}{\large\insertauthor\par}}\vspace*{1ex}
        \usebeamercolor{institute in headline}{\color{fg}{\large\insertinstitute\par}}     
      \end{column}
      \begin{column}{.02\paperwidth}
      \end{column}
    \end{columns}
    \vskip2ex
  \end{beamercolorbox}

}

%\tcbset{colframe=structure,fonttitle=\bfseries,beamer}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
\setbeamertemplate{footline}{
 
    \leavevmode%

  \begin{beamercolorbox}[wd=\paperwidth]{footline}
        \begin{columns}[T]
            \begin{column}{.02\paperwidth}
            \end{column}
            %
            \begin{column}{.95\paperwidth}
                \raggedleft
                \begin{column}{.25\paperwidth}
                  \begin{center}
                      \includegraphics[width=.45\linewidth]{logo.png}
                  \end{center}
            	\end{column}
            \end{column}
            %
            \begin{column}{.02\paperwidth}
            \end{column}
        \end{columns}
        \vskip2ex
     \end{beamercolorbox}

% % \begin{tcolorbox}[enhanced,watermark graphics=logo/onderkant,width=\paperwidth,
% % %watermark opacity=0.5,
% % watermark color=blue,watermark stretch=1]
% \begin{tcolorbox}[watermark graphics=logo/onderkant,width=\paperwidth,
% watermark opacity=1.0,%0.5,
% watermark stretch=1]
%         \begin{columns}[T]
%             \begin{column}{.02\paperwidth}
%             \end{column}
%             %
%             \begin{column}{.95\paperwidth}
%                 \vskip4ex
%                 \raggedleft
%                 \begin{column}{.25\paperwidth}
%                   \vskip8ex
%                   \begin{center}
%                       \includegraphics[width=.95\linewidth]{logo/ru_e_a4_wit_2014}
%                   \end{center}
%                   \vskip2ex
%             	\end{column}
%             \end{column}
%             %
%             \begin{column}{.02\paperwidth}
%             \end{column}
%         \end{columns}
%         \vskip2ex
% \end{tcolorbox}
}

%\beamertemplategridbackground[5cm]

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
\mode<all>

% adapt height of imtemize rectangles
\setbeamertemplate{itemize items}[triangle]
\setbeamertemplate{itemize item}{\raisebox{0.12ex}{$\blacktriangleright$}\hskip0.1em}
\setbeamertemplate{itemize subitem}{\raisebox{0.12ex}{$\triangleright$}\hskip0.1em}
% or define your own template using \defbeamertemplate{itemize item}, see beameruserguide.pdf

% equal font sizes for all levels
\setbeamerfont{itemize/enumerate body}{size=\large}
\setbeamerfont{itemize/enumerate subbody}{size=\large}
\setbeamerfont{itemize/enumerate subsubbody}{size=\large}











\makeatletter
\newcommand\beamerboxesframed[2][]{%
  \global\let\beamer@firstlineitemizeunskip=\relax%
  \vbox\bgroup%
  \setkeys{beamerboxes}{upper=block title,lower=block body,width=\textwidth}%
  \setkeys{beamerboxes}{#1}%
  {%
    \usebeamercolor{\bmb@lower}%
    \globalcolorstrue%
    \colorlet{lower.bg}{bg}%
  }%
  {%
    \usebeamercolor{\bmb@upper}%
    \globalcolorstrue%
    \colorlet{upper.bg}{bg}%
  }%
  %
  % Typeset head
  %
  \vskip4bp
  \setbox\bmb@box=\hbox{%
    \begin{minipage}[b]{\bmb@width}%
      \usebeamercolor[fg]{\bmb@upper}%
      #2%
    \end{minipage}}%
  \ifdim\wd\bmb@box=0pt%
    \setbox\bmb@box=\hbox{}%
    \ht\bmb@box=0pt%
    \bmb@prevheight=-4.5pt%
  \else%
    \wd\bmb@box=\bmb@width%
    \bmb@temp=\dp\bmb@box%
    \ifdim\bmb@temp<1.5pt%
      \bmb@temp=1.5pt%
    \fi%
    \setbox\bmb@box=\hbox{\raise\bmb@temp\hbox{\box\bmb@box}}%
    \dp\bmb@box=0pt%
    \bmb@prevheight=\ht\bmb@box%
  \fi%
  \bmb@temp=\bmb@width%
  \bmb@dima=\bmb@temp\advance\bmb@dima by2.2bp%
  \bmb@dimb=\bmb@temp\advance\bmb@dimb by4bp%
  \hbox{%
    \begin{pgfpicture}{0bp}{+-\ht\bmb@box}{0bp}{+-\ht\bmb@box}
      \ifdim\wd\bmb@box=0pt%
        \color{lower.bg}%
      \else%
        \color{upper.bg}%
      \fi%
      \pgfpathqmoveto{-4bp}{-1bp}
      \pgfpathqcurveto{-4bp}{1.2bp}{-2.2bp}{3bp}{0bp}{3bp}
      \pgfpathlineto{\pgfpoint{\bmb@temp}{3bp}}
      \pgfpathcurveto%
      {\pgfpoint{\bmb@dima}{3bp}}%
      {\pgfpoint{\bmb@dimb}{1.2bp}}%
      {\pgfpoint{\bmb@dimb}{-1bp}}%
      \bmb@dima=-\ht\bmb@box%
      \advance\bmb@dima by-2pt%
      \pgfpathlineto{\pgfpoint{\bmb@dimb}{\bmb@dima}}
      \pgfpathlineto{\pgfpoint{-4bp}{\bmb@dima}}
      \pgfpathclose
      \pgfsetstrokecolor{black}\pgfusepath{stroke, fill}
    \end{pgfpicture}%
    \copy\bmb@box%
  }%
  \nointerlineskip%
  \ifdim\wd\bmb@box=0pt
  \else
    \vskip2.4pt%
  \fi%
  \nointerlineskip%
  \setbox\bmb@colorbox=\hbox{{\pgfpicturetrue\pgfsetcolor{lower.bg}}}%
  \setbox\bmb@box=\hbox\bgroup\begin{minipage}[b]{\bmb@width}%
    \vskip2pt%
    \usebeamercolor[fg]{\bmb@lower}%
    \colorlet{beamerstructure}{upper.bg}%
    \colorlet{structure}{upper.bg}%
    %\color{.}%
}






\def\endbeamerboxesframed{%
  \end{minipage}\egroup%
  \wd\bmb@box=\bmb@width%
  \bmb@temp=\dp\bmb@box%
  \advance\bmb@temp by.5pt%
  \setbox\bmb@box=\hbox{\raise\bmb@temp\hbox{\box\bmb@box}}%
  \dp\bmb@box=0pt%
  \bmb@temp=\wd\bmb@box%
  \bmb@dima=\bmb@temp\advance\bmb@dima by2.2bp%
  \bmb@dimb=\bmb@temp\advance\bmb@dimb by4bp%
  \hbox{%
    \begin{pgfpicture}{0bp}{0bp}{0bp}{0bp}
      \unhbox\bmb@colorbox%
      \pgfpathmoveto{\pgfpoint{-4bp}{\ht\bmb@box}}
      \pgfpathlineto{\pgfpoint{-4bp}{1bp}}
      \pgfpathqcurveto{-4bp}{-1.2bp}{-2.2bp}{-3bp}{0bp}{-3bp}
      \pgfpathlineto{\pgfpoint{\the\bmb@temp}{-3bp}}
      \pgfpathcurveto%
      {\pgfpoint{\the\bmb@dima}{-3bp}}%
      {\pgfpoint{\the\bmb@dimb}{-1.2bp}}%
      {\pgfpoint{\the\bmb@dimb}{1bp}}%
      {
      \bmb@dima=\ht\bmb@box%
      \pgfpathlineto{\pgfpoint{\bmb@dimb}{\bmb@dima}}
      \pgfsetstrokecolor{black}\pgfusepath{stroke, fill}
      }
    \end{pgfpicture}%
    \box\bmb@box%
  }%
  \vskip2bp%
  \egroup% of \vbox\bgroup
}
\makeatother






\defbeamertemplateparent{blocks}[framed]{block begin,block end,%
  block alerted begin,block alerted end,%
  block example begin,block example end}[1][]
{[#1]}

\defbeamertemplate{block begin}{framed}[1][]
{
  \par\vskip\medskipamount%
  \begin{beamerboxesframed}[upper=block title,lower=block body,#1]%
    {\raggedright\usebeamerfont*{block title}\insertblocktitle}%
    \raggedright%
    \usebeamerfont{block body}%
}
\defbeamertemplate{block end}{framed}[1][]
{\end{beamerboxesframed}\vskip\smallskipamount}

\defbeamertemplate{block alerted begin}{framed}[1][]
{
  \par\vskip\medskipamount%
  \begin{beamerboxesframed}[upper=block title alerted,lower=block body alerted,#1]%
    {\raggedright\usebeamerfont*{block title alerted}\insertblocktitle}%
    \raggedright%
    \usebeamerfont{block body alerted}%
}%
\defbeamertemplate{block alerted end}{framed}[1][]
{\end{beamerboxesframed}\vskip\smallskipamount}

\defbeamertemplate{block example begin}{framed}[1][]
{
  \par\vskip\medskipamount%
  \begin{beamerboxesframed}[upper=block title example,lower=block body example,#1]
    {\raggedright\usebeamerfont*{block title example}\insertblocktitle}%
    \raggedright%
    \usebeamerfont{block body alerted}%
}%
\defbeamertemplate{block example end}{framed}[1][]
{\end{beamerboxesframed}\vskip\smallskipamount}





